package zapsetup

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// defaultCfg = newConfigDefault()
	defaultLog = NewLogger()
)

func RootLogger() *LoggerX {
	return defaultLog
}

func NewLogger(opts ...Option) *LoggerX {
	newCfg := newConfigDefault()
	for _, opt := range opts {
		opt(newCfg)
	}
	return newLoggerWrapper(newCfg)
}

func newConfigDefault() *zap.Config {
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
