package handle

import (
	"encoding/json"
	"time"
	"ton-event-idx/internal/app"
	"ton-event-idx/pkg/spreq"
	"ton-event-idx/pkg/tcntr"

	"github.com/sirupsen/logrus"
)

func Start(seqno int) {
	tr := &app.TONRPC

	seriGetBlockTrsns, err := tr.GetBlockTransactions(&tcntr.GetBlockTransactions{
		Workchain: app.CFG.StartBlockWorkchain,
		Shard:     app.CFG.StartBlockShard,
		Seqno:     seqno,
	})

	if err != nil {
		logrus.Fatal(err)
	}

	rawBlockResp, err := spreq.SendJsonPostReq(
		app.CFG.JsonRpcURL,
		app.CFG.JsonRpcTimeout,
		seriGetBlockTrsns,
	)

	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(string(rawBlockResp))

	var blockResp tcntr.BasicResp[tcntr.BlockTransactions]
	json.Unmarshal(rawBlockResp, &blockResp)

	logrus.Info(len(blockResp.Result), seqno, blockResp.OK)
	time.Sleep(time.Duration(app.CFG.PollingDelayMs) * time.Millisecond)
	Start(seqno + 1)
	// for _, v := range blockResp.Result {
	// 	logrus.Info(v)
	// }
}
