package heraldserver

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	discovery "github.com/lafetz/quest-micro/common/consul"
	commongrpc "github.com/lafetz/quest-micro/common/grpc"
	mailv1 "github.com/lafetz/quest-micro/proto/gen/mail/v1"
	herald "github.com/lafetz/quest-micro/services/herald/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HeraldServer struct {
	serviceName   string
	instanceID    string
	port          int
	HeraldService herald.HeraldServiceApi
	mailv1.UnimplementedEmailServiceServer
	logger   *slog.Logger
	registry discovery.RegistryApi
}

func (h *HeraldServer) Run() {
	grpcServer := commongrpc.NewServer(h.logger)
	mailv1.RegisterEmailServiceServer(grpcServer, h)
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(h.port))
	if err != nil {
		h.logger.Error("Failed to listen", "error", err)
		return
	}
	go func() {
		healthServer.SetServingStatus(h.serviceName, grpc_health_v1.HealthCheckResponse_SERVING)

		if err := h.registry.Register(h.instanceID, h.serviceName, fmt.Sprintf("%s:%d", h.serviceName, h.port)); err != nil {
			h.logger.Error("Failed to register Service", "error", err.Error())
			os.Exit(1)
		}
		if err := grpcServer.Serve(lis); err != nil {
			h.logger.Error("Grpc server failed", "error", err.Error())
			os.Exit(1)
		}
	}()

	h.gracefulStop(grpcServer, healthServer)
}

func (h *HeraldServer) gracefulStop(grpcServer *grpc.Server, healthServer *health.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	healthServer.SetServingStatus(h.serviceName, grpc_health_v1.HealthCheckResponse_NOT_SERVING)
	h.registry.Deregister(h.instanceID, h.serviceName)
	h.logger.Info("Shutting down server")
	grpcServer.GracefulStop()
	h.logger.Info("Server gracefully stopped")
}
func NewHeraldServer(serviceName string, instanceID string, registry discovery.RegistryApi, heraldService herald.HeraldServiceApi, port int, logger *slog.Logger) *HeraldServer {
	return &HeraldServer{
		serviceName:   serviceName,
		instanceID:    instanceID,
		HeraldService: heraldService,
		registry:      registry,
		port:          port,
		logger:        logger,
	}
}
