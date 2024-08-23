package main

import (
	"log/slog"
	"os"

	"github.com/lafetz/quest-micro/common/config"
	discovery "github.com/lafetz/quest-micro/common/consul"
	"github.com/lafetz/quest-micro/common/logger"
	herald "github.com/lafetz/quest-micro/services/herald/core"
	gomailadapter "github.com/lafetz/quest-micro/services/herald/gomail"
	heraldserver "github.com/lafetz/quest-micro/services/herald/server"
)

const (
	serviceName = "herald"
	version     = "1.0.0"
)

func main() {
	config, err := config.NewConfig(
		config.WithLogLevel(), config.WithPort(), config.WithRegistryURI(), config.WithEnv(),
		config.WithSMTPHost(),
		config.WithSMTPPort(),
		config.WithSMTPUsername(),
		config.WithSMTPPassword(),
	)

	instanceId := discovery.GenerateInstanceID(serviceName)
	if err != nil {
		slog.Error("config error", "error", err.Error(), "serviceName", serviceName, "instanceId", instanceId)
		os.Exit(1)
	}
	log := logger.NewLogger(config.Env, config.LogLevel, serviceName, version, instanceId)
	emailSender := gomailadapter.NewGomailSender(
		config.SMTPhost,
		config.SMTPport,
		config.SMTPusername,
		config.SMTPpassword,
	)
	svc := herald.NewHeraldService(emailSender)
	registry, err := discovery.NewConsulRegistry(config.RegistryURI)
	if err != nil {
		log.Error("unable to create new Registry", "error", err.Error())
		os.Exit(1)
	}
	srv := heraldserver.NewHeraldServer(serviceName, instanceId, registry, svc, config.Port, log)
	log.Info("Server running")
	srv.Run()

}
