package main

import (
	"os"

	"github.com/lafetz/quest-demo/common/logger"
	quest "github.com/lafetz/quest-demo/services/quest/core"
	"github.com/lafetz/quest-demo/services/quest/repository"
	"github.com/lafetz/quest-demo/services/quest/web"
	// "github.com/lafetz/quest-demo/services/knight/repository"
)

type MockKnightSrv struct{}

func (m *MockKnightSrv) GetKnightStatus(string) (bool, error) {
	return true, nil
}
func main() {
	log := logger.NewLogger("debug")
	db, err := repository.OpenDB("postgresql://user:password@postgres/quest?sslmode=disable")
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
