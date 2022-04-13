# 简介
本仓库旨在提供适用于一般场合的 zaplog 配置,做好"最后一公里".

# 繁介
zaplog 有着强大又灵活的配置项,但多数场合并不需要太多功能. 其自带的 `Production` 也仅适用在超高吞吐量的场合.
因此, 当你想关闭采样输出, 或者想在控制台查看日志, 又或者想要动态修改 `log.Level`, 还或者
想实现一些`sink`把日志投递到日志存储&分析系统里, 那就看看这个仓库里的代码吧.


# Usage
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