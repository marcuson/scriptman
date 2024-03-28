//go:build prod

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func getLoggerCore(logPath string) zapcore.Core {
	lumber := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    25, // megabytes
		MaxBackups: 3,
	}

	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(pe)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(lumber), zap.InfoLevel),
	)

	return core
}
