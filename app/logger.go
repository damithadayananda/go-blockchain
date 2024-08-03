package app

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"os"
)

var Logger logger

type logger struct {
	Error log.Logger
	Warn  log.Logger
	Info  log.Logger
	Trace log.Logger
}

func init() {
	baseLogger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	baseLogger = log.With(baseLogger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	logger := logger{
		Error: level.Error(baseLogger),
		Warn:  level.Warn(baseLogger),
		Info:  level.Info(baseLogger),
		Trace: level.Debug(baseLogger),
	}
	Logger = logger
}
