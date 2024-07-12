package main

import (
	knight "github.com/lafetz/quest-demo/services/knight/core"
	grpcserver "github.com/lafetz/quest-demo/services/knight/grpc"
	"github.com/lafetz/quest-demo/services/knight/repository"
)

func main() {
	mongo, err := repository.NewDb("mongodb://admin:admin11@localhost:27017/knight?authSource=admin")
	if err != nil {
		panic(err)
	}
	store, err := repository.NewStore(mongo)
	if err != nil {
		panic(err)
	}
	srv := knight.NewKnightService(store)
	grpc := grpcserver.NewGrpcServer(srv, 8080)
	grpc.Run()
}
