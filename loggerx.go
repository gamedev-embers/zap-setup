package zapsetup

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerX struct {
	*zap.Logger
	level zap.AtomicLevel
	logf  *zap.SugaredLogger
}

func newLoggerX(cfg *zap.Config) *LoggerX {
	log, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return &LoggerX{
		Logger: log,
		level:  cfg.Level,
		logf:   log.Sugar(),
	}
}

// SetLevel atomically
func (l *LoggerX) SetLevel(level zapcore.Level) *LoggerX {
	l.level.SetLevel(level)
	return l
}

func (l *LoggerX) WithSink(s Sink, logLevel ...zapcore.Level) *LoggerX {
	core := NewCoreX(l.Logger, s)
	if len(logLevel) > 0 {
		core.SetLogLevelForSink(logLevel[0])
	}
	l.Logger = zap.New(core, zap.WithCaller(true))
	return l
}

func (l *LoggerX) Debugf(format string, args ...interface{}) {
	l.logf.Debugf(format, args...)
}

func (l *LoggerX) Infof(format string, args ...interface{}) {
	l.logf.Infof(format, args...)
}

func (l *LoggerX) Warnf(format string, args ...interface{}) {
	l.logf.Warnf(format, args...)
}

func (l *LoggerX) Errorf(format string, args ...interface{}) {
	l.logf.Errorf(format, args...)
}
