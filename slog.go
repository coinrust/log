package log

import (
	"github.com/coinrust/slog"
)

type SLogger struct {
	inner *slog.Logger
}

func (l *SLogger) Debug(args ...interface{}) {
	slog.Debug(args...)
}

func (l *SLogger) Info(args ...interface{}) {
	slog.Info(args...)
}

func (l *SLogger) Warn(args ...interface{}) {
	slog.Warn(args...)
}

func (l *SLogger) Error(args ...interface{}) {
	slog.Error(args...)
}

func (l *SLogger) DPanic(args ...interface{}) {

}

func (l *SLogger) Panic(args ...interface{}) {
	slog.Crit(args...)
}

func (l *SLogger) Fatal(args ...interface{}) {
	slog.Crit(args...)
}

func (l *SLogger) Debugf(template string, args ...interface{}) {
	slog.Debugf(template, args...)
}

func (l *SLogger) Infof(template string, args ...interface{}) {
	slog.Infof(template, args...)
}

func (l *SLogger) Warnf(template string, args ...interface{}) {
	slog.Warnf(template, args...)
}

func (l *SLogger) Errorf(template string, args ...interface{}) {
	slog.Errorf(template, args...)
}

func (l *SLogger) DPanicf(template string, args ...interface{}) {
	slog.Critf(template, args...)
}

func (l *SLogger) Panicf(template string, args ...interface{}) {
	slog.Critf(template, args...)
}

func (l *SLogger) Fatalf(template string, args ...interface{}) {
	slog.Critf(template, args...)
}

func (l *SLogger) Debugw(msg string, keysAndValues ...interface{}) {

}

func (l *SLogger) Infow(msg string, keysAndValues ...interface{}) {

}

func (l *SLogger) Warnw(msg string, keysAndValues ...interface{}) {

}

func (l *SLogger) Errorw(msg string, keysAndValues ...interface{}) {

}

func (l *SLogger) DPanicw(msg string, keysAndValues ...interface{}) {

}

func (l *SLogger) Panicw(msg string, keysAndValues ...interface{}) {

}

func (l *SLogger) Fatalw(msg string, keysAndValues ...interface{}) {

}

func (l *SLogger) Sync() error {
	if l.inner == nil {
		return nil
	}
	l.inner.Close()
	return nil
}

func NewSLogger(c *Configuration) *SLogger {
	w, err := slog.NewTimedRotatingFileWriter(c.Path,
		slog.RotateByDate, 16)
	if err != nil {
		return nil
	}
	level := parseSLogLevel(c.Level)
	h := slog.NewHandler(level, slog.DefaultFormatter)
	h.AddWriter(w)
	l := slog.NewLogger(level)
	l.AddHandler(h)
	slog.SetDefaultLogger(l)

	return &SLogger{inner: l}
}

func parseSLogLevel(level string) slog.Level {
	var l slog.Level
	switch level {
	case DebugLevel:
		l = slog.DebugLevel
	case InfoLevel:
		l = slog.InfoLevel
	case WarnLevel:
		l = slog.WarnLevel
	case ErrorLevel:
		l = slog.ErrorLevel
	case PanicLevel:
		l = slog.CritLevel
	}
	return l
}
