package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		Level: zap.NewAtomicLevelAt(zap.InfoLevel),

		OutputPaths: []string{"stdout"},
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: nil,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
	var err error
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func GetLogger() *zap.Logger {
	return log
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	err := log.Sync()
	if err != nil {
		return
	}
}

func Error(msg string, err error, tags ...zap.Field) {
	if err != nil {
		tags = append(tags, zap.NamedError("error", err))
	}
	log.Error(msg, tags...)
	err = log.Sync()
	if err != nil {
		return
	}
}
