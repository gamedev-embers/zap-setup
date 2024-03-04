package main

import (
	"fmt"

	zapsetup "github.com/gamedev-embers/zap-setup"
	"github.com/gamedev-embers/zap-setup/sink/aliyunsls"
	"go.uber.org/zap"
)

var (
	log  = zapsetup.RootLogger()
	log2 = log.WithSink(aliyunSLS())
	log3 = log.WithSink(tencentCLS())

	// aliyunUrl  = flag.String("aliyun-sls", "", "aliyun sls url")
	// tencentUrl = flag.String("tencent-cls", "", "tencent cls url")
)

func main() {
	log.Debug("here is the default root logger")
	log.Info("here is the default root logger")
	log.Warn("here is the default root logger")
	log.Error("here is the default root logger")

	// log with aliyun sls
	log2.Warn("here is a logger with aliyun sls",
		zap.String("str", "value"),
		zap.Int("int", 1),
		zap.Error(fmt.Errorf("fake error")))

	// log with tencent cls
	log3.Error("here is a logger with tencentcloud cls",
		zap.String("str", "value"),
		zap.Int("int", 1),
		zap.Error(fmt.Errorf("fake error")))
}

func aliyunSLS() zapsetup.Sink {
	sink, err := aliyunsls.New("aliyun+sls://user:passwd@endpoint/projectA/logstoreA")
	if err != nil {
		panic(err)
	}
	sink.Open()
	return sink
}

func tencentCLS() zapsetup.Sink {
	sink, err := aliyunsls.New("tencent+cls://user:passwd@endpoint/projectA/logstoreA")
	if err != nil {
		panic(err)
	}
	sink.Open()
	return sink
}
