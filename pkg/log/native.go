package log

import (
	"log/slog"
	"os"
)

type nativeLog struct {
	slog *slog.Logger
}

func NewNativeLog() Log {
	return nativeLog{
		slog: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}

func (l nativeLog) Info(msg string, args ...any) {
	l.slog.Info(msg, args...)
}

func (l nativeLog) Debug(msg string, args ...any) {
	l.slog.Debug(msg, args...)
}

func (l nativeLog) Warn(msg string, args ...any) {
	l.slog.Warn(msg, args...)
}

func (l nativeLog) Error(msg string, args ...any) {
	l.slog.Error(msg, args...)
}
