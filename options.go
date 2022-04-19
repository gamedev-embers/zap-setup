package zapsetup

import (
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
