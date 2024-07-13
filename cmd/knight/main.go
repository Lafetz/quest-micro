package main

import (
	"os"

	"github.com/lafetz/quest-demo/common/logger"
	knight "github.com/lafetz/quest-demo/services/knight/core"
	grpcserver "github.com/lafetz/quest-demo/services/knight/grpc"
	"github.com/lafetz/quest-demo/services/knight/repository"
)

func main() {
	log := logger.NewLogger("debug")
	mongo, err := repository.NewDb("mongodb://admin:admin11@localhost:27017/knight?authSource=admin")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	store, err := repository.NewStore(mongo)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	srv := knight.NewKnightService(store)

	grpc := grpcserver.NewGrpcServer(srv, 8080, log)
	grpc.Run()
}
