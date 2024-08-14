package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/google/uuid"
	commongrpc "github.com/lafetz/quest-micro/common/grpc"
	quest "github.com/lafetz/quest-micro/quest/core"
	client "github.com/lafetz/quest-micro/quest/knight_grpc"
	web "github.com/lafetz/quest-micro/quest/server"
)

type MockQuestRepository struct{}

func (m *MockQuestRepository) AddQuest(ctx context.Context, quest quest.Quest) (*quest.Quest, error) {
	// Always returns the input quest without any error
	return &quest, nil
}

func (m *MockQuestRepository) GetAssignedQuests(ctx context.Context, knightID string) ([]*quest.Quest, error) {
	// Always returns an empty slice of Quests without any error
	return []*quest.Quest{}, nil
}

func (m *MockQuestRepository) CompleteQuest(ctx context.Context, questID uuid.UUID) error {
	// Always returns nil (no error)
	return nil
}

func (m *MockQuestRepository) GetQuest(ctx context.Context, questID uuid.UUID) (*quest.Quest, error) {
	// Always returns a dummy quest without any error
	return quest.NewQuest("xx", "ss", "xzzx", "xxxw"), nil
}
func main() {
	// config, err := configqst.NewConfig()
	// if err != nil {
	// 	slog.Error(err.Error())
	// 	os.Exit(1)
	// }
	log := slog.Default() //logger.NewLogger(config.Env, 0)
	// db, err := repository.OpenDB(config.DbUrl)
	// if err != nil {
	// 	log.Error(err.Error())
	// 	os.Exit(1)
	// }

	//		store := repository.NewDb(db)
	cb := commongrpc.NewCb("knightGrpc")
	grpcClient, err := client.NewKnightClient("localhost:8080", cb)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	log.Info("connected to grpc service...")
	kntSrv := client.NewKntClient(log, grpcClient)

	srv := quest.NewQuestService(&MockQuestRepository{}, kntSrv)
	app := web.NewApp(srv, 3000, log)
	err = app.Run()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
