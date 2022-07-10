package app

import (
	"context"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"ton-event-idx/pkg/psql"
	"unsafe"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/constraints"
)

type basicBlockInfo struct {
	WC    int32
	Shard uint64
	SeqNo uint32
}

type sleepInfo struct {
	MinDiff, MaxDiff, IfCantd int64
	MaxCount                  int
}

type mainConfig struct {
	LITE_SERVER_HOST string // ip:port
	LITE_SERVER_PKEY string // base64

	BlockInfo *basicBlockInfo
	SleepInfo *sleepInfo
	Database  *psql.PsqlConfig
}

func parseEnvInt[I constraints.Integer](varInt *I, envName string) {
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

func parseTime[I constraints.Integer](varTime *time.Duration, timeType time.Duration, envName string) {
	var envTime I
	parseEnvInt(&envTime, envName)

	*varTime = time.Duration(envTime) * timeType
}

var CFG mainConfig = mainConfig{
	BlockInfo: &basicBlockInfo{},
	SleepInfo: &sleepInfo{
		MinDiff:  10,  // ms
		MaxDiff:  100, // ms
		IfCantd:  100, // ms
		MaxCount: 10,  // slice size
	},
	Database: &psql.PsqlConfig{},
}

var DBCONN *pgxpool.Pool

func Configure() {
	logrus.Info("start \"Configure\" function")

	CFG.LITE_SERVER_HOST = os.Getenv("LITE_SERVER_HOST")
	CFG.LITE_SERVER_PKEY = os.Getenv("LITE_SERVER_PKEY")

	parseEnvInt(&CFG.BlockInfo.WC, "BLOCK_WC")
	parseEnvInt(&CFG.BlockInfo.SeqNo, "BLOCK_SEQNO")
	parseEnvInt(&CFG.BlockInfo.Shard, "BLOCK_SHARD")

	CFG.Database.Username = os.Getenv("DB_USER")
	CFG.Database.Password = os.Getenv("DB_PASW")
	CFG.Database.Database = os.Getenv("DB_NAME")
	CFG.Database.Host = os.Getenv("DB_HOST")
	CFG.Database.Port = os.Getenv("DB_PORT")

	parseEnvInt(&CFG.Database.MaxConnRetry, "DB_CONN_MAX_RETRY")
	parseEnvInt(&CFG.Database.RetryTimeout, "DB_CONN_TIMEOUT_S")

	parseTime[int](&CFG.Database.RetryTimeout, time.Second, "DB_CONN_TIMEOUT_S")
	parseTime[int](&CFG.Database.RetNeedSleep, time.Second, "DB_RETRY_SSLEEP_S")

	var err error
	DBCONN, err = psql.NewPsqlClient(context.Background(), CFG.Database)
	if err != nil {
		logrus.Fatal(err)
	}
}
