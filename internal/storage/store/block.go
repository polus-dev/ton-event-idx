package store

import (
	"context"
	"ton-event-idx/internal/storage/mcblock"
	"ton-event-idx/pkg/utils/dbutil"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type Event struct {
	Body []byte
}

type Block struct {
	WorkChain int32
	Shard     []byte
	SeqNo     uint32
	Events    []Event
}

func StoreBlock(
	ctx context.Context, tx pgx.Tx,
	mc *mcblock.MCBlockModel, blocks []Block,
) error {
	insertMcQ := `
		INSERT INTO mc_block (shard, seqno)
		VALUES ($1::bytea, $2) RETURNING id
	`

	insertBlockQ := `
		INSERT INTO block (mc_block_id, work_chain, shard, seqno)
		VALUES ($1::uuid, $2, $3::bytea, $4) RETURNING id
	`

	insertEventQ := `
		INSERT INTO event (block_id, contract_id, body)
		VALUES ($1::uuid, $2, $3::bytea) RETURNING id
	`

	err := tx.QueryRow(ctx, insertMcQ, mc.Shard, mc.SeqNo).Scan(&mc.ID)
	if err != nil {
		return dbutil.MapPgErr(err)
	}

	for _, block := range blocks {
		var blockId []byte

		if err := tx.QueryRow(
			ctx, insertBlockQ,
			mc.ID, block.WorkChain, block.Shard, block.SeqNo,
		).Scan(&blockId); err != nil {
			return dbutil.MapPgErr(err)
		}

		for _, event := range block.Events {
			contractId := 1 // TODO: ...
			var idd []byte
			if err := tx.QueryRow(
				ctx, insertEventQ,
				blockId, contractId, event.Body,
			).Scan(&idd); err != nil {
				return dbutil.MapPgErr(err)
			}

		}
	}

	if err := tx.Commit(ctx); err != nil {
		logrus.Fatal(err)
		return err
	}

	return nil
}
