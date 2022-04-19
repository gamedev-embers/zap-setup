package zapsetup

import (
	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"
)

type CoreX struct {
	core zapcore.Core
	sink Sink
}

func NewCoreX(l *zap.Logger, sink Sink) *CoreX {
	return &CoreX{
		core: l.Core(),
		sink: sink,
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
	c.sink.Write(ent, fields)
	return c.core.Write(ent, fields)
}

func (c *CoreX) Sync() error {
	return c.core.Sync()
}
