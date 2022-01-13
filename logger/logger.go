package logger

import (
	"github.com/kataras/golog"
)

// Logger logger
type Logger interface {
	Fatal(v ...interface{})
	Fatalf(format string, args ...interface{})
	Error(v ...interface{})
	Errorf(format string, args ...interface{})
	Warn(v ...interface{})
	Warnf(format string, args ...interface{})
	Info(v ...interface{})
	Infof(format string, args ...interface{})
	Debug(v ...interface{})
	Debugf(format string, args ...interface{})

	Child(interface{}) *golog.Logger
	SetPrefix(s string) *golog.Logger
	Clone() *golog.Logger
}
