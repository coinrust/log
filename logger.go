package log

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	PanicLevel = "panic"
)

var logger Logger

type Logger interface {
	Debug(args ...interface{})

	Info(args ...interface{})

	Warn(args ...interface{})

	Error(args ...interface{})

	//DPanic(args ...interface{})

	Panic(args ...interface{})

	//Fatal(args ...interface{})

	Debugf(template string, args ...interface{})

	Infof(template string, args ...interface{})

	Warnf(template string, args ...interface{})

	Errorf(template string, args ...interface{})

	//DPanicf(template string, args ...interface{})

	Panicf(template string, args ...interface{})

	//Fatalf(template string, args ...interface{})

	//Debugw(msg string, keysAndValues ...interface{})

	//Infow(msg string, keysAndValues ...interface{})

	//Warnw(msg string, keysAndValues ...interface{})

	//Errorw(msg string, keysAndValues ...interface{})

	//DPanicw(msg string, keysAndValues ...interface{})

	//Panicw(msg string, keysAndValues ...interface{})

	//Fatalw(msg string, keysAndValues ...interface{})

	Sync() error
}

// Init init logger
func Init(path, level string, options ...Option) {
	c := &Configuration{
		Path:        path,
		Level:       level,
		MaxFileSize: 512, //MB
		MaxBackups:  100,
		MaxAge:      60,
		Compress:    true,
		Caller:      false,
	}

	for _, opt := range options {
		opt(c)
	}

	switch {
	case c.SLog:
		logger = NewSLogger(c)
	default:
		logger = NewZapLogger(c)
	}
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

// Sync flushes buffer, if any
func Sync() {
	if logger == nil {
		return
	}

	logger.Sync()
}
