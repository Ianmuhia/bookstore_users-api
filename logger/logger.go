package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log *zap.Logger
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
			EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: nil,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
	var err error
	if Log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}
