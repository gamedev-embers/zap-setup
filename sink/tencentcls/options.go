package tencentcls

import (
	"os"

	"go.uber.org/zap/zapcore"
)

var LEVELS = map[string]zapcore.Level{
	"panic":  zapcore.PanicLevel,
	"dpanic": zapcore.DPanicLevel,
	"error":  zapcore.ErrorLevel,
	"warn":   zapcore.WarnLevel,
	"info":   zapcore.InfoLevel,
	"debug":  zapcore.DebugLevel,
}

type Option func(*Sink)

// WithLogLevel: "debug" "info", "warn", "error"
func WithLogLevel(l string) Option {
	logLevel := LEVELS[l]
	return func(s *Sink) {
		s.logLevel = logLevel
	}
}

// Aliyun SLS setups
func WithProject(project, logstore, source string) Option {
	return func(s *Sink) {
		s.project = project
		s.logstore = logstore
		s.source = source
	}
}

// Aliyun SLS source field
func WithHostSource() Option {
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return func(s *Sink) {
		s.source = host
	}
}
