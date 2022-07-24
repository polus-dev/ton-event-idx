package scan

import (
	"context"
	"time"
	"ton-event-idx/internal/app"
	"ton-event-idx/internal/storage/mcblock"
	"ton-event-idx/pkg/utils/mmath"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"github.com/xssnick/tonutils-go/ton"
)

func maxTime(x, y time.Duration) time.Duration {
	if x < y {
		return y
	}
	return x
}

func process(mc *mcblock.MCBlockDTO, api *ton.APIClient) error {
	lastOk := time.Now()
	timeDiffs := []time.Duration{0}

	for {
		block, err := api.LookupBlock(
			context.Background(),
			app.MASTER_CHAIN_ID, mc.Shard, mc.SeqNo,
		)

		if err != nil {
			// TODO: fix infinity loop (mark block as can't be processed)
			logrus.Infof("sleep for %d milliseconds", app.CFG.SleepInfo.IfCantd.Milliseconds())
			time.Sleep(app.CFG.SleepInfo.IfCantd)
			continue
		}

		mc.SeqNo += 1
		sleepFor := mmath.CalcAvgInteger(timeDiffs)

		timeDiffs = append(
			timeDiffs, maxTime(
				app.CFG.SleepInfo.MinDiff,
				time.Since(lastOk)-app.CFG.SleepInfo.MaxDiff,
			),
		)

		tx, err := app.DBCONN.BeginTx(
			context.Background(),
			pgx.TxOptions{IsoLevel: pgx.Serializable},
		)

		if err != nil {
			return err
		}

		if err := processMcBlock(tx, block, api); err != nil {
			logrus.Fatal(err)
			tx.Rollback(context.Background())
		}

		logrus.Infof("sleep for %d milliseconds", sleepFor.Milliseconds())
		time.Sleep(sleepFor)

		lastOk = time.Now()
		if len(timeDiffs) > app.CFG.SleepInfo.MaxCount {
			timeDiffs = timeDiffs[1:]
		}
	}
}

func StartScanMasterChain(api *ton.APIClient) error {
	blockRepo := mcblock.NewRepo(app.DBCONN)

	isEmpty, err := blockRepo.IsEmpty(context.Background())
	if err != nil {
		return err
	}

	if isEmpty {
		return process(app.CFG.BlockInfo, api)
	}

	lastBlock, err := blockRepo.SelectLast(context.Background())
	if err != nil {
		return err
	}

	return process(lastBlock.ToDTO(), api)
}
