package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/lafetz/quest-micro/common/config"
	discovery "github.com/lafetz/quest-micro/common/consul"
	commongrpc "github.com/lafetz/quest-micro/common/grpc"
	"github.com/lafetz/quest-micro/common/logger"
	httpgrpc "github.com/lafetz/quest-micro/http-grpc"
)

const (
	serviceName = "grpcGateway"
	grpcService = "knight"
	version     = "1.0.0"
)

func main() {

	config, err := config.NewConfig(config.WithPort(), config.WithLogLevel(), config.WithRegistryURI(), config.WithEnv())
	instanceId := discovery.GenerateInstanceID(serviceName)
	if err != nil {
		slog.Error("config error", "error", err.Error(), "serviceName", serviceName, "instanceId", instanceId)
		os.Exit(1)
	}
	log := logger.NewLogger(config.Env, config.LogLevel, serviceName, version, instanceId)
	registry, err := discovery.NewConsulRegistry(config.RegistryURI)

	if err != nil {
		log.Error("failed to create new Registry", "error", err.Error())
		os.Exit(1)
	}

	addrs, err := registry.ServiceAddresses(grpcService)
	if err != nil {
		log.Error("failed to get service address", "error", err.Error())
		os.Exit(1)
	}
	cb := commongrpc.NewCb("knightGrpc")
	conn, err := commongrpc.NewGRPCClient(addrs[0], cb)
	if err != nil {
		log.Error("unable to make connection with grpc service", "error", err.Error())
		os.Exit(1)
	}
	srv, err := httpgrpc.NewGatewayServer(conn, config.Port, slog.Default())
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	log.Info("Server running")
	if err := srv.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to listen and serve", err)
	}

}
