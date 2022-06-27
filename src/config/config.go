package config

import (
	"os"
	"strconv"

	rus "github.com/sirupsen/logrus"
)

type mainConfig struct {
	JsonRpcURL          string
	PollingDelayMs      int
	StartBlockSeqno     int
	StartBlockShard     int
	StartBlockWorkchain int
}

var CFG mainConfig
var LOG *rus.Logger = rus.New()

func parseEnvInt(varInt *int, envName string) {
	parsed, err := strconv.Atoi(os.Getenv(envName))
	if err != nil {
		LOG.Errorf("can't parse %s from env", envName)
	}

	*varInt = parsed
}

func Configure() {
	LOG = &rus.Logger{
		Out:   os.Stderr,
		Hooks: make(rus.LevelHooks),
		Level: rus.DebugLevel,
		Formatter: &rus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	}

	CFG.JsonRpcURL = os.Getenv("JSON_PRC")
	parseEnvInt(&CFG.PollingDelayMs, "POLLING_DELAY_MS")

	parseEnvInt(&CFG.StartBlockSeqno, "START_BLOCK_SEQNO")
	parseEnvInt(&CFG.StartBlockShard, "START_BLOCK_SHARD")
	parseEnvInt(&CFG.StartBlockWorkchain, "START_BLOCK_WORKCHAIN")
}
