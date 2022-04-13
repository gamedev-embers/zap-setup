# 简介
本项目旨在提供适用于一般场合的 zaplog 配置, 搞定"最后一公里".  
官方提供的默认配置包含采样输出,数字时间等等,仅适用在超高吞吐量的场合.


# 使用
```go
import (
	zapsetup "github.com/upgrade-or-die/zap-setup"
)

var (
	log  = zapsetup.RootLogger()
	log2 = log.Sugar()
)

func main() {
	log.Info("hello zaplog")
	log2.Infof("hello %s", "小明")

	// change log level on the fly
	log.SetLevel(zap.DebugLevel)
	log.Debug("hello zaplog")
	log2.Infof("hello %s", "小明")
}
```

# TODO 
- [ ] aliyun-sls-sink
- [ ] aws-watchlog-sink

# LICENCE
MIT License