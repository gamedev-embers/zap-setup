package tencentcls

import (
	"fmt"
	"strconv"
	"time"

	cls "github.com/tencentcloud/tencentcloud-cls-sdk-go"

	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/proto"
)

func fields2logs(ent zapcore.Entry, fields []zapcore.Field) *cls.Log {
	var headerSize int
	if ent.Stack == "" {
		headerSize = 4
	} else {
		headerSize = 5
	}
	y, m, d := ent.Time.Date()

	contents := make([]*cls.Log_Content, headerSize+len(fields))
	contents[0] = &cls.Log_Content{
		Key:   proto.String("date"),
		Value: proto.String(fmt.Sprintf("%04d-%02d-%02d", y, int(m), d)),
	}
	contents[1] = &cls.Log_Content{
		Key:   proto.String("msg"),
		Value: proto.String(ent.Message),
	}
	contents[2] = &cls.Log_Content{
		Key:   proto.String("level"),
		Value: proto.String(ent.Level.String()),
	}
	contents[3] = &cls.Log_Content{
		Key:   proto.String("caller"),
		Value: proto.String(ent.Caller.TrimmedPath()),
	}
	if ent.Stack != "" {
		contents[4] = &cls.Log_Content{
			Key:   proto.String("stacktrace"),
			Value: proto.String(ent.Stack),
		}
	}

	for i := 0; i < len(fields); i++ {
		var f = &fields[i]
		var key = f.Key
		var val string
		switch f.Type {
		case zapcore.StringType:
			val = f.String
		case zapcore.Int64Type, zapcore.Int32Type, zapcore.Int16Type, zapcore.Int8Type:
			val = strconv.FormatInt(f.Integer, 10)
		case zapcore.Uint64Type, zapcore.Uint32Type, zapcore.Uint16Type, zapcore.Uint8Type:
			val = strconv.FormatUint(uint64(f.Integer), 10)
		case zapcore.DurationType:
			val = strconv.FormatInt(f.Integer/int64(time.Millisecond), 10)
		case zapcore.ErrorType:
			val = fmt.Sprintf("%v", f.Interface)
		case zapcore.TimeType:
			if f.Interface == nil {
				val = time.Unix(0, f.Integer).Format(time.RFC3339)
			} else {
				tz := f.Interface.(*time.Location)
				val = time.Unix(0, f.Integer).In(tz).Format(time.RFC3339)
			}
		case zapcore.ReflectType:
			if f.Interface == nil {
				val = "null"
			} else {
				val = fmt.Sprintf("%+v", f.Interface)
			}
		default:
			val = fmt.Sprintf("%v", f.Interface)
		}
		contents[i+headerSize] = &cls.Log_Content{
			Key:   proto.String(key),
			Value: proto.String(val),
		}
	}
	return &cls.Log{
		Time:     proto.Int64(ent.Time.Unix()),
		Contents: contents,
	}
}
