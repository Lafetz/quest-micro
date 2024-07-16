package main

import (
	"log"
	"os"

	"github.com/lafetz/quest-demo/common/logger"
	configqst "github.com/lafetz/quest-demo/quest/config"
	quest "github.com/lafetz/quest-demo/quest/core"
	"github.com/lafetz/quest-demo/quest/repository"
	"github.com/lafetz/quest-demo/quest/web"
)

type MockKnightSrv struct{}

func (m *MockKnightSrv) GetKnightStatus(string) (bool, error) {
	return true, nil
}
func main() {
	config, err := configqst.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	log := logger.NewLogger(config.Env)
	db, err := repository.OpenDB(config.DbUrl)
	store := repository.NewDb(db)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	mockK := &MockKnightSrv{}
	srv := quest.NewQuestService(store, mockK)
	app := web.NewApp(srv, 3000, log)
	err = app.Run()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

//""
