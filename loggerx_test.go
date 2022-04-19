package zapsetup

import (
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

func TestLoggerX_withSink(t *testing.T) {
	assert := assert.New(t)
	sink := newSinkForTest()
	log := NewLogger().WithSink(sink)

	fields := []zap.Field{zap.Int("int", 1), zap.String("str", "2")}
	for i := 0; i < 10; i++ {
		msg := strconv.Itoa(i)
		log.Info(msg, fields...)
		select {
		case row := <-sink.queue:
			assert.Equal(msg, row.ent.Message)
			assert.Equal(fields, row.fields)
		case <-time.After(1 * time.Second):
			assert.FailNow("sink write failed")
		}
	}
}
