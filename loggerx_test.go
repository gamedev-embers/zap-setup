package zapsetup

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type sinkForTest struct {
	queue chan struct {
		ent    zapcore.Entry
		fields []zapcore.Field
	}
}

func newSinkForTest() *sinkForTest {
	return &sinkForTest{
		queue: make(chan struct {
			ent    zapcore.Entry
			fields []zapcore.Field
		}, 16),
	}
}
func (s *sinkForTest) Open()  {}
func (s *sinkForTest) Close() {}
func (s *sinkForTest) Write(ent zapcore.Entry, fields []zapcore.Field) {
	s.queue <- struct {
		ent    zapcore.Entry
		fields []zapcore.Field
	}{ent, fields}
}

func TestLoggerX_WithSink_All(t *testing.T) {
	assert := assert.New(t)
	sink := newSinkForTest()
	opt := WithLogLevel(zapcore.DebugLevel)
	log := NewLogger(opt).WithSink(sink)

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("msg %d", i)
		fields := []zap.Field{zap.Int("int", i), zap.String("str", strconv.Itoa(i))}
		log.Debug(msg+" debug", fields...)
		log.Info(msg+" info", fields...)
		log.Warn(msg+" warn", fields...)
		log.Error(msg+" error", fields...)

		select {
		case row := <-sink.queue:
			assert.Equal(msg+" debug", row.ent.Message)
			assert.Equal(fields, row.fields)
		case <-time.After(100 * time.Millisecond):
			assert.FailNow("sink write failed")
		}

		select {
		case row := <-sink.queue:
			assert.Equal(msg+" info", row.ent.Message)
			assert.Equal(fields, row.fields)
		case <-time.After(100 * time.Millisecond):
			assert.FailNow("sink write failed")
		}

		select {
		case row := <-sink.queue:
			assert.Equal(msg+" warn", row.ent.Message)
			assert.Equal(fields, row.fields)
		case <-time.After(100 * time.Millisecond):
			assert.FailNow("sink write failed")
		}

		select {
		case row := <-sink.queue:
			assert.Equal(msg+" error", row.ent.Message)
			assert.Equal(fields, row.fields)
		case <-time.After(100 * time.Millisecond):
			assert.FailNow("sink write failed")
		}
	}
}

func TestLoggerX_WithSink_DisableInfoLevel(t *testing.T) {
	assert := assert.New(t)
	sink := newSinkForTest()
	log := NewLogger().WithSink(sink, zapcore.WarnLevel)

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("msg %d", i)
		fields := []zap.Field{zap.Int("int", i), zap.String("str", strconv.Itoa(i))}
		log.Debug(msg+" debug", fields...)
		log.Info(msg+" info", fields...)
		log.Warn(msg+" warn", fields...)
		log.Error(msg+" error", fields...)
		log.Infof(msg + " info2")
		select {
		case row := <-sink.queue:
			assert.Equal(msg+" warn", row.ent.Message)
			assert.Equal(fields, row.fields)
		case <-time.After(100 * time.Millisecond):
			assert.FailNow("sink write failed")
		}

		select {
		case row := <-sink.queue:
			assert.Equal(msg+" error", row.ent.Message)
			assert.Equal(fields, row.fields)
		case <-time.After(100 * time.Millisecond):
			assert.FailNow("sink write failed")
		}
	}
}
