package main

import (
	"log/slog"
	"os"

	config "github.com/lafetz/quest-micro/common/config"
	discovery "github.com/lafetz/quest-micro/common/consul"
	"github.com/lafetz/quest-micro/common/logger"
	knight "github.com/lafetz/quest-micro/services/knight/core"
	"github.com/lafetz/quest-micro/services/knight/repository"
	knightserver "github.com/lafetz/quest-micro/services/knight/server"
)

const (
	serviceName = "knight"
	version     = "1.0.0"
)

func main() {

	config, err := config.NewConfig(config.WithLogLevel(), config.WithPort(), config.WithDbUrl(), config.WithRegistryURI(), config.WithEnv())
	instanceId := discovery.GenerateInstanceID(serviceName)
	if err != nil {
		slog.Error("config error", "error", err.Error(), "serviceName", serviceName, "instanceId", instanceId)
		os.Exit(1)
	}
	log := logger.NewLogger(config.Env, config.LogLevel, serviceName, version, instanceId)
	mongo, close, err := repository.NewDb(config.DbUrl, log)
	defer close()
	if err != nil {
		log.Error("failed to connect with db", "error", err.Error())
		os.Exit(1)
	}
	store, err := repository.NewStore(mongo)
	if err != nil {
		log.Error("failed to create Store", "error", err.Error())
		os.Exit(1)

	}
	svc := knight.NewKnightService(store)
	registry, err := discovery.NewConsulRegistry(config.RegistryURI)
	if err != nil {
		log.Error("unable to create new Registry", "error", err.Error())
		os.Exit(1)
	}

	srv := knightserver.NewKnightServer(serviceName, instanceId, registry, svc, config.Port, log)
	log.Info("Server running")
	srv.Run()
}
