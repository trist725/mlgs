package db

import (
	mlog "github.com/trist725/mgsu/log"
)

var gLogger = mlog.NewConsoleLogger()

func Logger() mlog.ILogger {
	return gLogger
}

func SetLogger(l mlog.ILogger) {
	gLogger = l
}
