package logger

import (
	"log/slog"
	"os"
)

var LogLevels = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func NewLogger(level string) *slog.Logger {

	logLevel := LogLevels[level]

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	})

	logger := slog.New(logHandler)
	return logger
}
