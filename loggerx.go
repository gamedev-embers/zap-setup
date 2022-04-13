package zapsetup

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerX struct {
	*zap.Logger
	level zap.AtomicLevel
}

func newLoggerWrapper(cfg *zap.Config) *LoggerX {
	log, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return &LoggerX{
		Logger: log,
		level:  cfg.Level,
	}
}

// SetLevel atomically
func (l *LoggerX) SetLevel(level zapcore.Level) {
	l.level.SetLevel(level)
}
