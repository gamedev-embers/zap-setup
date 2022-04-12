package zapsetup

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option func(*zap.Config)

func WithSampling(c *zap.SamplingConfig) Option {
	return func(cfg *zap.Config) {
		cfg.Sampling = c
	}
}

func WithLogLevel(level zapcore.Level) Option {
	return func(cfg *zap.Config) {
		cfg.Level.SetLevel(level)
	}
}

func WithAliyunSLS() Option {
	panic(fmt.Errorf("not implemented yet"))
	return func(cfg *zap.Config) {
		// TODO: implement aliyun+sls://region-id.xxx.xxx/projectId/logstore?accessId=xx&ccessKey=yy
	}
}

func WithAWSWatchLog() Option {
	panic(fmt.Errorf("not implemented yet"))
	return func(cfg *zap.Config) {
		// TODO: implement aws+watchlog://region-id.xxx.xxx/projectId/logstore?accessId=xx&ccessKey=yy
	}
}
