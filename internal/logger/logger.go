package logger

import (
	"github.com/adrg/xdg"
	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
)

func init() {
	logPath, _ := xdg.DataFile("scriptman/cli.log")
	core := getLoggerCore(logPath)
	l := zap.New(core)
	logger = l.Sugar()
}

func Get() *zap.SugaredLogger {
	return logger
}
