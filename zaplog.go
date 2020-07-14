package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// ZapLogger 默认会使用zap作为日志输出引擎. Log集成了日志切割的功能。默认文件大小512M，自动压缩
type ZapLogger struct {
	*Configuration

	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

// createLumberjackHook 创建LumberjackHook，其作用是为了将日志文件切割，压缩
func (l *ZapLogger) createLumberjackHook() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   l.Path,
		MaxSize:    l.MaxFileSize,
		MaxBackups: l.MaxBackups,
		MaxAge:     l.MaxAge,
		Compress:   l.Compress,
	}
}

func (l *ZapLogger) Build() {
	writeSyncer := []zapcore.WriteSyncer{
		zapcore.AddSync(l.createLumberjackHook()),
	}

	if l.Stdout {
		writeSyncer = append(writeSyncer, zapcore.Lock(os.Stdout))
	}

	var level zapcore.Level
	switch l.Level {
	case DebugLevel:
		level = zap.DebugLevel
	case InfoLevel:
		level = zap.InfoLevel
	case WarnLevel:
		level = zap.WarnLevel
	case ErrorLevel:
		level = zap.ErrorLevel
	case PanicLevel:
		level = zap.PanicLevel
	default:
		level = zap.InfoLevel
	}

	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = zapcore.ISO8601TimeEncoder
	//cnf := zapcore.NewJSONEncoder(conf)
	cnf := zapcore.NewConsoleEncoder(conf)
	core := zapcore.NewCore(cnf,
		zapcore.NewMultiWriteSyncer(writeSyncer...),
		level)

	l.logger = zap.New(core)
	if l.Caller {
		l.logger = l.logger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(2))
	}
	l.sugar = l.logger.Sugar()
}

func (l *ZapLogger) Debug(args ...interface{}) {
	l.sugar.Debug(args...)
}

func (l *ZapLogger) Info(args ...interface{}) {
	l.sugar.Info(args...)
}

func (l *ZapLogger) Warn(args ...interface{}) {
	l.sugar.Warn(args...)
}

func (l *ZapLogger) Error(args ...interface{}) {
	l.sugar.Error(args...)
}

func (l *ZapLogger) DPanic(args ...interface{}) {
	l.sugar.DPanic(args...)
}

func (l *ZapLogger) Panic(args ...interface{}) {
	l.sugar.Panic(args...)
}

func (l *ZapLogger) Fatal(args ...interface{}) {
	l.sugar.Fatal(args...)
}

func (l *ZapLogger) Debugf(template string, args ...interface{}) {
	l.sugar.Debugf(template, args...)
}

func (l *ZapLogger) Infof(template string, args ...interface{}) {
	l.sugar.Infof(template, args...)
}

func (l *ZapLogger) Warnf(template string, args ...interface{}) {
	l.sugar.Warnf(template, args...)
}

func (l *ZapLogger) Errorf(template string, args ...interface{}) {
	l.sugar.Errorf(template, args...)
}

func (l *ZapLogger) DPanicf(template string, args ...interface{}) {
	l.sugar.DPanicf(template, args...)
}

func (l *ZapLogger) Panicf(template string, args ...interface{}) {
	l.sugar.Panicf(template, args...)
}

func (l *ZapLogger) Fatalf(template string, args ...interface{}) {
	l.sugar.Fatalf(template, args...)
}

func (l *ZapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.sugar.Debugw(msg, keysAndValues...)
}

func (l *ZapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.sugar.Infow(msg, keysAndValues...)
}

func (l *ZapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.sugar.Warnw(msg, keysAndValues...)
}

func (l *ZapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.sugar.Errorw(msg, keysAndValues...)
}

func (l *ZapLogger) DPanicw(msg string, keysAndValues ...interface{}) {
	l.sugar.DPanicw(msg, keysAndValues...)
}

func (l *ZapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	l.sugar.Panicw(msg, keysAndValues...)
}

func (l *ZapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.sugar.Fatalw(msg, keysAndValues...)
}

func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}

// NewZapLogger new ZapLogger
func NewZapLogger(c *Configuration) *ZapLogger {
	logger := &ZapLogger{Configuration: c}
	logger.Build()

	return logger
}
