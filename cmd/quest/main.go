package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"github.com/lafetz/quest-micro/common/config"
	discovery "github.com/lafetz/quest-micro/common/consul"
	commongrpc "github.com/lafetz/quest-micro/common/grpc"
	"github.com/lafetz/quest-micro/common/logger"
	quest "github.com/lafetz/quest-micro/services/quest/core"
	client "github.com/lafetz/quest-micro/services/quest/knight_grpc"
	web "github.com/lafetz/quest-micro/services/quest/server"
)

const (
	serviceName = "quest"
	grpcService = "knight"
	version     = "1.0.0"
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
	config, err := config.NewConfig(config.WithPort(), config.WithLogLevel(), config.WithRegistryURI(), config.WithEnv(), config.WithDbUrl())
	instanceId := discovery.GenerateInstanceID(serviceName)
	if err != nil {
		slog.Error("config error", "error", err.Error(), "serviceName", serviceName, "instanceId", instanceId)
		os.Exit(1)
	}
	log := logger.NewLogger(config.Env, config.LogLevel, serviceName, version, instanceId)
	registry, err := discovery.NewConsulRegistry(config.RegistryURI)

	addrs, err := registry.ServiceAddresses(grpcService)
	if err != nil {
		log.Error("failed to get service address", "error", err.Error())
		os.Exit(1)
	}
	if err != nil {
		log.Error("failed to create new Registry", "error", err.Error())
		os.Exit(1)
	}

	// db, err := repository.OpenDB(config.DbUrl)
	// if err != nil {
	// 	log.Error(err.Error())
	// 	os.Exit(1)
	// }

	//		store := repository.NewDb(db)
	cb := commongrpc.NewCb(grpcService)

	grpcClient, err := client.NewKnightClient(addrs[0], cb)
	if err != nil {
		log.Error("unable to make connection with grpc service", "error", err.Error())
		os.Exit(1)
	}

	kntSrv := client.NewKntClient(log, grpcClient)

	srv := quest.NewQuestService(&MockQuestRepository{}, kntSrv)
	app := web.NewApp(srv, 3000, log)
	err = app.Run()
	log.Info("Server running")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
