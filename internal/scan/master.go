package scan

import (
	"context"
	"time"
	"ton-event-idx/internal/app"
	"ton-event-idx/pkg/utils/mmath"

	"github.com/sirupsen/logrus"
	"github.com/xssnick/tonutils-go/ton"
)

func StartScanMasterChain(api *ton.APIClient) error {
	mc := *app.CFG.BlockInfo

	lastOk := time.Now()
	timeDiffs := []int64{0}

	for {

		block, err := api.LookupBlock(
			context.Background(),
			mc.WC, mc.Shard, mc.SeqNo,
		)

		if err != nil {
			// TODO: fix infinity loop (mark block as can't be processed)
			logrus.Infof("sleep for %d milliseconds", app.CFG.SleepInfo.IfCantd)
			time.Sleep(time.Duration(app.CFG.SleepInfo.IfCantd) * time.Millisecond)
			continue
		}

		go processMcBlock(block)
		mc.SeqNo += 1

		sleepFor := mmath.CalcAvgInteger(timeDiffs)

		logrus.Infof("sleep for %d milliseconds", sleepFor)
		time.Sleep(time.Duration(sleepFor) * time.Millisecond)

		sinceLast := time.Since(lastOk).Milliseconds()
		timeDiffs = append(
			timeDiffs,
			mmath.MaxInteger(
				app.CFG.SleepInfo.MinDiff,
				sinceLast-app.CFG.SleepInfo.MaxDiff,
			),
		)

		lastOk = time.Now()

		if len(timeDiffs) > app.CFG.SleepInfo.MaxCount {
			timeDiffs = timeDiffs[1:]
		}
	}
}
