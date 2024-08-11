package logger

import (
	"log/slog"
	"os"
)

func NewLogger(env string, lvl slog.Level, serviceName string, version string, instanceId string) *slog.Logger {

	var logHandler slog.Handler
	switch env {
	case "dev":
		logHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: lvl,
		})
	default:
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: lvl,
		})
	}

	logger := slog.New(logHandler).With(slog.Group(
		"service_info",
		slog.String("env", env),
		slog.String("service_Name", serviceName),
		slog.String("version", version),
		slog.String("instanceId", instanceId),
	),
	)

	return logger
}
