package tencentcls

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestUtils_fields2logs(t *testing.T) {
	assert := assert.New(t)
	assert.NoError(nil)

	enc := zapcore.Entry{
		Time:    time.Date(2022, 04, 01, 0, 0, 0, 0, time.Local),
		Message: "hello",
		Caller: zapcore.EntryCaller{
			Defined: true,
			File:    "/path/to/the/file",
			Line:    9527,
		},
	}
	fields := []zap.Field{
		zap.Int("int", 1),
		zap.String("str", "value"),
	}
	t.Run("no stacktrace", func(t *testing.T) {
		enc.Stack = ""
		L := func(k, v string) *sls.LogContent {
			return &sls.LogContent{
				Key:   &k,
				Value: &v,
			}
		}

		l := fields2logs(enc, fields)
		assert.Equal(4+2, len(l.Contents))
		assert.Equal(L("date", "2022-04-01"), l.Contents[0])
		assert.Equal(L("msg", "hello"), l.Contents[1])
		assert.Equal(L("level", "info"), l.Contents[2])
		assert.Equal(L("caller", "the/file:9527"), l.Contents[3])
		assert.Equal(L("int", "1"), l.Contents[4])
		assert.Equal(L("str", "value"), l.Contents[5])
		assert.Equal(uint32(enc.Time.Unix()), *l.Time)
	})

	t.Run("stacktrace", func(t *testing.T) {
		enc.Stack = "this is test traces"
		L := func(k, v string) *sls.LogContent {
			return &sls.LogContent{
				Key:   &k,
				Value: &v,
			}
		}

		l := fields2logs(enc, fields)
		assert.Equal(4+2+1, len(l.Contents))
		assert.Equal(L("date", "2022-04-01"), l.Contents[0])
		assert.Equal(L("msg", "hello"), l.Contents[1])
		assert.Equal(L("level", "info"), l.Contents[2])
		assert.Equal(L("caller", "the/file:9527"), l.Contents[3])
		assert.Equal(L("stacktrace", "this is test traces"), l.Contents[4])
		assert.Equal(L("int", "1"), l.Contents[5])
		assert.Equal(L("str", "value"), l.Contents[6])
		assert.Equal(uint32(enc.Time.Unix()), *l.Time)
	})

}

func BenchmarkUtils_fields2logs(b *testing.B) {
	enc := zapcore.Entry{
		Time:    time.Date(2022, 04, 01, 0, 0, 0, 0, time.Local),
		Message: "hello",
		Caller: zapcore.EntryCaller{
			Defined: true,
			File:    "/path/to/the/file",
			Line:    9527,
		},
	}
	buildFields := func(size int, typeName string) []zap.Field {
		rs := make([]zap.Field, size)
		for i := 0; i < size; i++ {
			key := fmt.Sprintf("key%d", i)
			switch typeName {
			case "int":
				rs[i] = zap.Int(key, i)
			case "str":
				rs[i] = zap.String(key, fmt.Sprintf("value-%d", i))
			default:
				panic(fmt.Errorf("invalid type %s", typeName))
			}
		}
		return rs
	}

	for _, size := range []int{1, 3, 5, 10, 30} {
		for _, typeName := range []string{"int", "str"} {
			name := fmt.Sprintf("%s/fields=%d", typeName, size)
			fields := buildFields(size, typeName)
			runtime.GC()
			b.Run(name, func(b *testing.B) {
				for i := 0; i <= b.N; i++ {
					l := fields2logs(enc, fields)
					if l == nil {
						b.Fatalf("nil result")
					}
					if len(l.Contents)-4 != size {
						b.Fatalf("invalid size")
					}
				}
			})
		}
	}
}
