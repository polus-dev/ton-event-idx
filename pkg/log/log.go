package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func init() {
	logrus.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		ForceColors:            true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(formatter)
}
