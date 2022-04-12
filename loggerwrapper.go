package zapsetup

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerWrapper struct {
	*zap.Logger
	level zap.AtomicLevel
}

func newLoggerWrapper(cfg *zap.Config) *LoggerWrapper {
	log, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return &LoggerWrapper{
		Logger: log,
		level:  cfg.Level,
	}
}

// SetLevel atomically
func (l *LoggerWrapper) SetLevel(level zapcore.Level) {
	l.level.SetLevel(level)
}
