# 简介
本仓库旨在提供适用于一般场合的 zaplog 配置,做好"最后一公里".

# 繁介
zaplog 有着强大又灵活的配置项,但多数场合并不需要太多功能. 其自带的 `Production` 也仅适用在超高吞吐量的场合.
因此, 当你想关闭采样输出, 或者想在控制台查看日志, 又或者想要动态修改 `log.Level`, 还或者
想实现一些`sink`把日志投递到日志存储&分析系统里, 那就看看这个仓库里的代码吧.

# 特性
* 开箱即用
* 支持阿里云日志(aliyun-sls)
* 支持腾讯云日志(tencent-cls)

# Usage
```go
import (
	"flag"
	"fmt"

	zapsetup "github.com/gamedev-embers/zap-setup"
	"github.com/gamedev-embers/zap-setup/sink/aliyunsls"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log  = zapsetup.NewLogger()
)

func init() {
	log = log.WithSink(aliyunSLS(), zapcore.WarnLevel)
	// log = log.WithSink(tencentCLS(), zapcore.WarnLevel)
}

func main() {
	log.Debug("here is the default root logger")
	log.Info("here is the default root logger")

	// log with aliyun sls
	log.Warn("here is a logger with aliyun sls",
		zap.String("str", "value"),
		zap.Int("int", 1),
		zap.Error(fmt.Errorf("fake error")))

	log.Error("here is a logger with aliyun sls",
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


```

# TODO 
- [x] aliyun-sls-sink
- [x] tencent-cls-sink
- [ ] aws-watchlog-sink

# LICENCE
MIT License
