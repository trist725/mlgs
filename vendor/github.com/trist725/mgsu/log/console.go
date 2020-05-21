package log

import (
	"fmt"
	l "log"

	"github.com/astaxie/beego/logs"
)

type ConsoleLogger struct {
	sign  string
	level int
}

func NewConsoleLogger() (l ILogger) {
	l = &ConsoleLogger{
		level: logs.LevelDebug,
	}
	return
}

func NewConsoleLoggerWithSign(sign string) (l ILogger) {
	l = &ConsoleLogger{
		sign:  sign,
		level: logs.LevelDebug,
	}
	return
}

func (cl ConsoleLogger) Debug(format string, args ...interface{}) {
	if LevelDebug > cl.level {
		return
	}
	l.Printf("[D] %s\n", fmt.Sprintf(format, args...))
}

func (cl ConsoleLogger) Info(format string, args ...interface{}) {
	if LevelInformational > cl.level {
		return
	}
	l.Printf("[I] %s\n", fmt.Sprintf(format, args...))
}

func (cl ConsoleLogger) Warn(format string, args ...interface{}) {
	if LevelWarning > cl.level {
		return
	}
	l.Printf("[W] %s\n", fmt.Sprintf(format, args...))
}

func (cl ConsoleLogger) Error(format string, args ...interface{}) {
	if LevelError > cl.level {
		return
	}
	l.Printf("[E] %s\n", fmt.Sprintf(format, args...))
}

func (cl *ConsoleLogger) SetLevel(level int) {
	cl.level = level
}

func (ConsoleLogger) Close() {

}
