package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type mainConfig struct {
	JsonRpcURL          string
	PollingDelayMs      int
	StartBlockSeqno     int
	StartBlockShard     int
	StartBlockWorkchain int
}

var CFG mainConfig

func parseEnvInt(varInt *int, envName string) {
	parsed, err := strconv.Atoi(os.Getenv(envName))
	if err != nil {
		logrus.Errorf("can't parse %s from env", envName)
	}

	*varInt = parsed
}

func Configure() {
	CFG.JsonRpcURL = os.Getenv("JSON_PRC")
	parseEnvInt(&CFG.PollingDelayMs, "POLLING_DELAY_MS")

	parseEnvInt(&CFG.StartBlockSeqno, "START_BLOCK_SEQNO")
	parseEnvInt(&CFG.StartBlockShard, "START_BLOCK_SHARD")
	parseEnvInt(&CFG.StartBlockWorkchain, "START_BLOCK_WORKCHAIN")
}
