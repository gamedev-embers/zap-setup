package zapsetup

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerX struct {
	*zap.Logger
	level zap.AtomicLevel
}

func newLoggerX(cfg *zap.Config) *LoggerX {
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
func (l *LoggerX) SetLevel(level zapcore.Level) *LoggerX {
	l.level.SetLevel(level)
	return l
}

func (l *LoggerX) WithSink(s Sink) *LoggerX {
	core := NewCoreX(l.Logger, s)
	l.Logger = zap.New(core, zap.AddCaller())
	return l
}
