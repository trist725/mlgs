package log

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/trist725/mgsu/util"
)

const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

type ILogger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	SetLevel(int)
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	Close()
}

func New(logDir string, logFileBaseName string) ILogger {
	util.MustMkdirIfNotExist(logDir)

	var logger = logs.NewLogger(10000)

	logger.Async()

	//logger.EnableFuncCallDepth(true)

	//logger.SetLogFuncCallDepth(3)

	config := fmt.Sprintf(`{"filename":"%s/%s.log","level":%d,"maxlines":50000,"separate":["error"]}`,
		logDir, logFileBaseName, logs.LevelDebug)

	logger.SetLogger("multifile", config)

	return logger
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
var logger = NewConsoleLogger()

func SetLogger(l ILogger) {
	logger = l
}

func SetLogLevel(level int) {
	logger.SetLevel(level)
}

func Debug(format string, args ...interface{}) {
	logger.Debug("%s", fmt.Sprintf(format, args...))
}

func Info(format string, args ...interface{}) {
	logger.Info("%s", fmt.Sprintf(format, args...))
}

func Warn(format string, args ...interface{}) {
	logger.Warn("%s", fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
	logger.Error("%s", fmt.Sprintf(format, args...))
}
