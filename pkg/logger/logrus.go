package logger

import (
	"github.com/sirupsen/logrus"
)

type logrusLog struct {
}

func NewLogrus() Logger {
	return logrusLog{}
}

func (l logrusLog) Debug(msg ...any) {
	logrus.Debug(msg...)
}

func (l logrusLog) Debugf(format string, args ...any) {
	logrus.Debugf(format, args...)
}

func (l logrusLog) Info(msg ...any) {
	logrus.Info(msg...)
}

func (l logrusLog) Infof(format string, args ...any) {
	logrus.Infof(format, args...)
}

func (l logrusLog) Warn(msg ...any) {
	logrus.Warn(msg...)
}

func (l logrusLog) Warnf(format string, args ...any) {
	logrus.Warnf(format, args...)
}

func (l logrusLog) Error(msg ...any) {
	logrus.Error(msg...)
}

func (l logrusLog) Errorf(format string, args ...any) {
	logrus.Errorf(format, args...)
}
