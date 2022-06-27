package logger

import (
	"ton-event-idx/src/config"
)

func Info(args ...interface{}) {
	config.LOG.Info(args)
}

func Error(args ...interface{}) {
	config.LOG.Error(args)
}

func Debug(args ...interface{}) {
	config.LOG.Debug(args)
}

func Infof(format string, args ...interface{}) {
	config.LOG.Infof(format, args)
}

func Errorf(format string, args ...interface{}) {
	config.LOG.Errorf(format, args)
}

func Debugf(format string, args ...interface{}) {
	config.LOG.Debugf(format, args)
}
