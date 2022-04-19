package aliyunsls

import (
	"fmt"
	"sync"

	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"go.uber.org/zap/zapcore"
)

type Sink struct {
	producer *producer.Producer
	mux      sync.Mutex
	logLevel zapcore.Level

	project  string
	logstore string
	source   string
}

func New(_url string, opts ...Option) (*Sink, error) {
	u, err := ParseURL(_url)
	if err != nil {
		return nil, fmt.Errorf("invalid aliyunsls-url: %w", err)
	}
	cfg := producer.GetDefaultProducerConfig()
	cfg.Endpoint = u.Endpoint
	cfg.AccessKeyID = u.AccessKeyID
	cfg.AccessKeySecret = u.AccessKeySecret
	_producer := producer.InitProducer(cfg)
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
	s.producer.SafeClose()
	s.producer = nil
}

func (s *Sink) Write(ent zapcore.Entry, fields []zapcore.Field) {
	if s.logLevel >= ent.Level {
		return
	}
	l := fields2logs(ent, fields)
	topic := ent.Level.String()
	s.producer.SendLog(s.project, s.logstore, topic, s.source, l)
}

func (s *Sink) GetProducer() *producer.Producer {
	return s.producer
}
