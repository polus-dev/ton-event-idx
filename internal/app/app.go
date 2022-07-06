package app

import (
	"os"
	"strconv"
	"time"
	"ton-event-idx/pkg/tcntr"

	"github.com/sirupsen/logrus"
)

type mainConfig struct {
	JsonRpcURL          string
	PollingDelayMs      int
	StartBlockSeqno     int
	StartBlockShard     int
	StartBlockWorkchain int

	JsonRpcTimeout time.Duration
}

var CFG mainConfig
var TONRPC tcntr.TonCenterSerializer

func parseEnvInt(varInt *int, envName string) {
	parsed, err := strconv.Atoi(os.Getenv(envName))
	if err != nil {
		logrus.Errorf("can't parse %s from env", envName)
	}

	*varInt = parsed
}

func init() {
	logrus.Info("start \"Configure\" function")
	CFG.JsonRpcURL = os.Getenv("JSON_PRC")

	parseEnvInt(&CFG.PollingDelayMs, "POLLING_DELAY_MS")
	parseEnvInt(&CFG.StartBlockSeqno, "START_BLOCK_SEQNO")
	parseEnvInt(&CFG.StartBlockShard, "START_BLOCK_SHARD")
	parseEnvInt(&CFG.StartBlockWorkchain, "START_BLOCK_WORKCHAIN")

	CFG.JsonRpcTimeout = 1 * time.Second
	TONRPC = tcntr.TonCenterSerializer{JsonRpcURL: CFG.JsonRpcURL}
}
