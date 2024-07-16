package logger

import (
	"log/slog"
	"os"
)

var LogLevels = map[string]slog.Level{
	"dev":  slog.LevelDebug,
	"prod": slog.LevelWarn,
}

func NewLogger(env string) *slog.Logger {
	logLevel := LogLevels[env]
	var logHandler slog.Handler
	switch env {
	case "dev":
		logHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     logLevel,
		})
	default:
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     logLevel,
		})
	}

	logger := slog.New(logHandler)
	return logger
}
