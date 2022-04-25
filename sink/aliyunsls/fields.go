package aliyunsls

import (
	"fmt"
	"strconv"
	"time"

	sls "github.com/aliyun/aliyun-log-go-sdk"

	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/proto"
)

func fields2logs(ent zapcore.Entry, fields []zapcore.Field) *sls.Log {
	var headerSize int
	if ent.Stack == "" {
		headerSize = 4
	} else {
		headerSize = 5
	}
	contents := make([]*sls.LogContent, headerSize+len(fields))

	y, m, d := ent.Time.Date()
	contents[0] = &sls.LogContent{
		Key:   proto.String("date"),
		Value: proto.String(fmt.Sprintf("%04d-%02d-%02d", y, int(m), d)),
	}
	contents[1] = &sls.LogContent{
		Key:   proto.String("msg"),
		Value: proto.String(ent.Message),
	}
	contents[2] = &sls.LogContent{
		Key:   proto.String("level"),
		Value: proto.String(ent.Level.String()),
	}
	contents[3] = &sls.LogContent{
		Key:   proto.String("caller"),
		Value: proto.String(ent.Caller.TrimmedPath()),
	}
	if ent.Stack != "" {
		contents[4] = &sls.LogContent{
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
			val = time.Duration(f.Integer).String()
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
		contents[i+headerSize] = &sls.LogContent{
			Key:   proto.String(key),
			Value: proto.String(val),
		}
	}

	row := &sls.Log{
		Time:     proto.Uint32(uint32(ent.Time.Unix())),
		Contents: contents,
	}
	return row
}
