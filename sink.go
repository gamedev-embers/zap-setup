package zapsetup

import "go.uber.org/zap/zapcore"

type Sink interface {
	Open()
	Close()
	Write(ent zapcore.Entry, fields []zapcore.Field)
}
