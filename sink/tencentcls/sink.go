package tencentcls

import (
	"fmt"
	"sync"

	cls "github.com/tencentcloud/tencentcloud-cls-sdk-go"
	"go.uber.org/zap/zapcore"
)

type Sink struct {
	producer *cls.AsyncProducerClient
	mux      sync.Mutex
	logLevel zapcore.Level

	project  string
	logstore string
	source   string
}

func New(_url string, opts ...Option) (*Sink, error) {
	u, err := ParseURL(_url)
	if err != nil {
		return nil, fmt.Errorf("invalid tencentcls-url: %w", err)
	}

	cfg := cls.GetDefaultAsyncProducerClientConfig()
	cfg.Endpoint = u.Endpoint
	cfg.AccessKeyID = u.AccessKeyID
	cfg.AccessKeySecret = u.AccessKeySecret
	_producer, err := cls.NewAsyncProducerClient(cfg)
	if err != nil {
		return nil, err
	}

	sink := &Sink{
		producer: _producer,
		project:  u.Project,
		logstore: u.LogStore,
		source:   "default",
	}
	for _, opt := range opts {
		opt(sink)
	}
	return sink, nil
}

func (s *Sink) Open() {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.producer == nil {
		panic(fmt.Errorf("nil producer"))
	}
	s.producer.Start()
}

func (s *Sink) Close() {
	if s.producer == nil {
		return
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.producer == nil {
		return
	}
	s.producer.Close(60 * 1000)
	s.producer = nil
}

func (s *Sink) Write(ent zapcore.Entry, fields []zapcore.Field) {
	if s.logLevel >= ent.Level {
		return
	}
	l := fields2logs(ent, fields)
	topic := ent.Level.String()

	cb := Callback{}
	s.producer.SendLog(topic, l, &cb)
}

func (s *Sink) GetProducer() *cls.AsyncProducerClient {
	return s.producer
}
