package log

import (
	"fmt"
)

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	PanicLevel = "panic"
)

var logger *Logger

// Init init logger
func Init(path, level string, options ...Option) {
	logger = NewLogger(path, level, options...)
}

// NewLogger new logger
func NewLogger(path, level string, options ...Option) *Logger {
	logger := NewZapAdapter(fmt.Sprintf("%s", path), level)

	for _, opt := range options {
		opt(logger)
	}

	logger.Build()

	return logger
}

// Sync flushes buffer, if any
func Sync() {
	if logger == nil {
		return
	}

	logger.logger.Sync()
}

// Debug 使用方法：log.Debug("test")
func Debug(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.Debug(args...)
}

// Debugf 使用方法：log.Debugf("test:%s", err)
func Debugf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.Debugf(template, args...)
}

// Debugw 使用方法：log.Debugw("test", "field1", "value1", "field2", "value2")
func Debugw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.Errorw(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.Panicw(msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.Fatalw(msg, keysAndValues...)
}
