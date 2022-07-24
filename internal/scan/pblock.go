package scan

import (
	"context"
	"encoding/binary"
	"ton-event-idx/internal/storage/mcblock"
	"ton-event-idx/internal/storage/store"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient/tlb"
	"github.com/xssnick/tonutils-go/ton"
)

func processMcBlock(tx pgx.Tx, masterBlock *tlb.BlockInfo, api *ton.APIClient) error {
	logrus.Infof("new masterchain block: %d", masterBlock.SeqNo)

	shards, err := api.GetBlockShardsInfo(context.TODO(), masterBlock)
	if err != nil {
		logrus.Warn("Error while GetBlockShardsInfo: ", err.Error())
		return err
	}

	var processed []store.Block

	for _, s := range shards {
		shardBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(shardBytes, s.Shard)

		var (
			block store.Block = store.Block{
				WorkChain: s.Workchain,
				Shard:     shardBytes,
				SeqNo:     s.SeqNo,
			}
			fetched []*tlb.TransactionID
			after   *tlb.TransactionID
			more    bool = true
		)

		for more {
			fetched, more, err = api.GetBlockTransactions(context.Background(), s, 100, after)
			if err != nil {
				logrus.Warn("Error while GetBlockTransactions: ", err.Error())
				return err
			}

			if more {
				after = fetched[len(fetched)-1]
			}

			for _, id := range fetched {
				tx, err := api.GetTransaction(
					context.Background(), s,
					address.NewAddress(0, 0, id.AccountID), // TODO: workchain from s
					id.LT,
				)

				if err != nil {
					logrus.Warn("Error while GetTransaction: ", err.Error())
					return err
				}

				for _, out := range tx.IO.Out {
					block.Events = append(block.Events, store.Event{
						Body: out.Msg.Payload().ToBOC(),
					})
				}

			}
		}
		processed = append(processed, block)
	}

	masterShardBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(masterShardBytes, masterBlock.Shard)

	err = store.StoreBlock(
		context.Background(), tx,
		&mcblock.MCBlockModel{
			Shard: masterShardBytes,
			SeqNo: masterBlock.SeqNo,
		},
		processed,
	)

	if err != nil {
		return err
	}

	return nil
}
