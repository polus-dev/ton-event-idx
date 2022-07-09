package app

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/sirupsen/logrus"
)

type basicBlockInfo struct {
	WC    int32
	Shard uint64
	SeqNo uint32
}

type sleepInfo struct {
	MinDiffSleep int
	MaxDiffSleep int
	IfCantGetDat int
	MaxDiffCount int
}

type mainConfig struct {
	LITE_SERVER_HOST string // ip:port
	LITE_SERVER_PKEY string // base64

	BlockInfo *basicBlockInfo
	SleepInfo *sleepInfo
}

var CFG mainConfig = mainConfig{
	BlockInfo: &basicBlockInfo{},
	SleepInfo: &sleepInfo{
		MinDiffSleep: 10,  // ms
		MaxDiffSleep: 100, // ms
		IfCantGetDat: 100, // ms
		MaxDiffCount: 10,  // slice size
	},
}

func parseEnvInt[I int32 | uint32 | uint64](varInt *I, envName string) {
	intSize := int(unsafe.Sizeof(*varInt) * 8) // convert uintptr to int

	strType := reflect.TypeOf(varInt).String()
	var parsed I

	if strings.Contains(strType, "u") {
		p, err := strconv.ParseUint(os.Getenv(envName), 10, intSize)
		if err != nil {
			logrus.Fatalf("can't ParseUint \"%s\" env; err: %s", envName, err)
		}

		parsed = I(p)
	} else {
		p, err := strconv.ParseInt(os.Getenv(envName), 10, intSize)
		if err != nil {
			logrus.Fatalf("can't ParseInt \"%s\" env; err: %s", envName, err)
		}

		parsed = I(p)
	}

	*varInt = parsed
}

func Configure() {
	logrus.Info("start \"Configure\" function")

	CFG.LITE_SERVER_HOST = os.Getenv("LITE_SERVER_HOST")
	CFG.LITE_SERVER_PKEY = os.Getenv("LITE_SERVER_PKEY")

	parseEnvInt(&CFG.BlockInfo.WC, "BLOCK_WC")
	parseEnvInt(&CFG.BlockInfo.SeqNo, "BLOCK_SEQNO")
	parseEnvInt(&CFG.BlockInfo.Shard, "BLOCK_SHARD")
}
