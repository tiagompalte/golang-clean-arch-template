package logger

type Logger interface {
	Debug(msg ...any)
	Debugf(format string, args ...any)
	Info(msg ...any)
	Infof(format string, args ...any)
	Warn(msg ...any)
	Warnf(format string, args ...any)
	Error(msg ...any)
	Errorf(format string, args ...any)
}
