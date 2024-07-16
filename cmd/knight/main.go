package main

import (
	"log"

	"github.com/lafetz/quest-demo/common/logger"
	configKnt "github.com/lafetz/quest-demo/knight/config"
	knight "github.com/lafetz/quest-demo/knight/core"
	grpcserver "github.com/lafetz/quest-demo/knight/grpc"
	"github.com/lafetz/quest-demo/knight/repository"
)

func main() {

	config, err := configKnt.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	logger := logger.NewLogger(config.Env)
	mongo, err := repository.NewDb(config.DbUrl)
	if err != nil {
		log.Fatal(err)

	}

	store, err := repository.NewStore(mongo)
	if err != nil {
		log.Fatal(err)

	}
	srv := knight.NewKnightService(store)

	grpc := grpcserver.NewGrpcServer(srv, 8080, logger)
	grpc.Run()
}
