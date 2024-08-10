package main

import (
	"log/slog"
	"os"

	"github.com/lafetz/quest-micro/common/logger"
	configKnt "github.com/lafetz/quest-micro/knight/config"
	knight "github.com/lafetz/quest-micro/knight/core"
	"github.com/lafetz/quest-micro/knight/repository"
	grpcserver "github.com/lafetz/quest-micro/knight/server"
)

func main() {

	config, err := configKnt.NewConfig()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	log := logger.NewLogger(config.Env)
	mongo, close, err := repository.NewDb(config.DbUrl, log)
	defer close()
	if err != nil {

		log.Error(err.Error())
		os.Exit(1)
	}
	log.Info("connected to DB")
	store, err := repository.NewStore(mongo)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)

	}
	srv := knight.NewKnightService(store)

	grpc := grpcserver.NewGrpcServer(srv, config.Port, log)
	log.Info("Server starting")
	grpc.Run()

}
