package main

import (
	"encoding/json"
	"fmt"
	"time"
	"ton-event-idx/src/config"
	"ton-event-idx/src/logger"
	"ton-event-idx/src/tonrpc"
)

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}

func main() {
	config.Configure()
	logger.Info("Starting the \"ton-event-idx\"")

	rpc := tonrpc.TonRPC{JsonRpcURL: config.CFG.JsonRpcURL}

	// resp, err := rpc.GetBlockTransactions(tonrpc.GetBlockTransactions{
	// 	Workchain: 0,
	// 	Shard:     8000000000000000,
	// 	Seqno:     26822299,
	// })

	// if err != nil {
	// 	logger.Error(err.Error())
	// }

	// var respDes tonrpc.BasicResp[tonrpc.BlockTransactions]
	// json.Unmarshal(resp, &respDes)

	// logger.Info(respDes.Result.Transactions)

	resp, err := rpc.GetTransactions(tonrpc.GetTransactions{
		Address: "0:a5b51fcf4cbbb036db5eefbb31f705d79ec118fa27cac0dcd893e1585029eaad",
		Hash:    "SPAN00z1yQf5rY/ihBgd8pcaAtmntcE+7YKo4vIRSSw=",
		LT:      29021823000003,
		Limit:   1,
	})

	if err != nil {
		logger.Error(err.Error())
	}

	fmt.Println(prettyPrint(resp))
	for {
		fmt.Println("ABOfA")
		time.Sleep(1 * time.Second)
	}
}
