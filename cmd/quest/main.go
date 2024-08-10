package main

import (
	"log/slog"
	"os"

	"github.com/lafetz/quest-micro/common/logger"
	configqst "github.com/lafetz/quest-micro/quest/config"
	quest "github.com/lafetz/quest-micro/quest/core"
	client "github.com/lafetz/quest-micro/quest/knight_grpc"
	"github.com/lafetz/quest-micro/quest/repository"
	web "github.com/lafetz/quest-micro/quest/server"
)

type MockKnightSrv struct{}

func (m *MockKnightSrv) GetKnightStatus(string) (bool, error) {
	return true, nil
}
func main() {
	config, err := configqst.NewConfig()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	log := logger.NewLogger(config.Env)
	db, err := repository.OpenDB(config.DbUrl)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	log.Info("connected to DB...")
	store := repository.NewDb(db)

	grpcClient, err := client.NewKnightClient("localhost:8080")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	log.Info("connected to grpc service...")
	kntSrv := client.NewKntClient(log, grpcClient)

	srv := quest.NewQuestService(store, kntSrv)
	app := web.NewApp(srv, 3000, log)
	err = app.Run()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
