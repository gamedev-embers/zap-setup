package zapsetup

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	rootLogger = NewLogger()
)

func RootLogger() *LoggerX {
	return rootLogger
}

func NewLogger(opts ...Option) *LoggerX {
	newCfg := defaultConfig()
	for _, opt := range opts {
		opt(newCfg)
	}
	return newLoggerX(newCfg)
}

func defaultConfig() *zap.Config {
	return &zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development:      false,
		Sampling:         nil,
		Encoding:         "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
}
