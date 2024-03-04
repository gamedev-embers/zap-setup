package zapsetup

import (
	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"
)

type CoreX struct {
	core zapcore.Core
	sink Sink

	logLevelForSink zapcore.Level
}

func NewCoreX(l *zap.Logger, sink Sink) *CoreX {
	return &CoreX{
		core:            l.Core(),
		sink:            sink,
		logLevelForSink: zapcore.DebugLevel,
	}
}

func (c *CoreX) Enabled(level zapcore.Level) bool {
	return c.core.Enabled(level)
}

func (c *CoreX) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *CoreX) With(fields []zapcore.Field) zapcore.Core {
	return c.core.With(fields)
}

func (c *CoreX) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	if c.logLevelForSink <= ent.Level {
		c.sink.Write(ent, fields)
		return nil
	}
	return c.core.Write(ent, fields)
}

func (c *CoreX) Sync() error {
	return c.core.Sync()
}

// SetLogLevelForSink doesn't change the log level of the core,
// but it changes the log level of the sink
// Note: it's not goroutine safe
func (c *CoreX) SetLogLevelForSink(level zapcore.Level) {
	c.logLevelForSink = level
}
