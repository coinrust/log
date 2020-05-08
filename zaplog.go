package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// Logger 默认会使用zap作为日志输出引擎. Log集成了日志切割的功能。默认文件大小512M，自动压缩

type Logger struct {
	Path        string // 文件路径，如：./app.log
	Level       string // 日志输出的级别
	MaxFileSize int    // 日志文件大小的最大值，单位(M)
	MaxBackups  int    // 最多保留备份数
	MaxAge      int    // 日志文件保存的时间，单位(天)
	Compress    bool   // 是否压缩
	Caller      bool   // 日志是否需要显示调用位置
	Stdout      bool   // 是否输出到控制台

	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

type Option func(log *Logger)

func SetStdout(b bool) Option {
	return func(log *Logger) {
		log.Stdout = b
	}
}

func SetMaxFileSize(size int) Option {
	return func(log *Logger) {
		log.MaxFileSize = size
	}
}

func SetMaxBackups(n int) Option {
	return func(log *Logger) {
		log.MaxBackups = n
	}
}

func SetMaxAge(age int) Option {
	return func(log *Logger) {
		log.MaxAge = age
	}
}

func SetCompress(compress bool) Option {
	return func(log *Logger) {
		log.Compress = compress
	}
}

func SetCaller(caller bool) Option {
	return func(log *Logger) {
		log.Caller = caller
	}
}

func (l *Logger) isCaller(level string) bool {
	return logger.Caller
}

func (l *Logger) maxFileSize(level string) int {
	return l.MaxFileSize
}

func (l *Logger) maxBackups(level string) int {
	return l.MaxBackups
}

func (l *Logger) maxAge(level string) int {
	return l.MaxAge
}

// createLumberjackHook 创建LumberjackHook，其作用是为了将日志文件切割，压缩
func (l *Logger) createLumberjackHook() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   l.Path,
		MaxSize:    l.MaxFileSize,
		MaxBackups: l.MaxBackups,
		MaxAge:     l.MaxAge,
		Compress:   l.Compress,
	}
}

func (l *Logger) Build() {
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

func (l *Logger) Debug(args ...interface{}) {
	l.sugar.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.sugar.Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.sugar.Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.sugar.Error(args...)
}

func (l *Logger) DPanic(args ...interface{}) {
	l.sugar.DPanic(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.sugar.Panic(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.sugar.Fatal(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.sugar.Debugf(template, args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.sugar.Infof(template, args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.sugar.Warnf(template, args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.sugar.Errorf(template, args...)
}

func (l *Logger) DPanicf(template string, args ...interface{}) {
	l.sugar.DPanicf(template, args...)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.sugar.Panicf(template, args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.sugar.Fatalf(template, args...)
}

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.sugar.Debugw(msg, keysAndValues...)
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.sugar.Infow(msg, keysAndValues...)
}

func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.sugar.Warnw(msg, keysAndValues...)
}

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.sugar.Errorw(msg, keysAndValues...)
}

func (l *Logger) DPanicw(msg string, keysAndValues ...interface{}) {
	l.sugar.DPanicw(msg, keysAndValues...)
}

func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.sugar.Panicw(msg, keysAndValues...)
}

func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.sugar.Fatalw(msg, keysAndValues...)
}

func NewZapAdapter(path, level string) *Logger {
	return &Logger{
		Path:        path,
		Level:       level,
		MaxFileSize: 512, //MB
		MaxBackups:  100,
		MaxAge:      60,
		Compress:    true,
		Caller:      false,
	}
}
