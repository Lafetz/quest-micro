package knightserver

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	discovery "github.com/lafetz/quest-micro/common/consul"
	commongrpc "github.com/lafetz/quest-micro/common/grpc"
	protoknight "github.com/lafetz/quest-micro/proto/gen"
	knight "github.com/lafetz/quest-micro/services/knight/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type KnightServer struct {
	serviceName   string //service name
	instanceID    string
	knightService knight.KnightServiceApi
	port          int
	protoknight.UnimplementedKnightServiceServer
	logger   *slog.Logger
	registry discovery.RegistryApi
}

func (g *KnightServer) Run() {

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(g.port))
	if err != nil {
		g.logger.Error("Failed to listen", "error", err)
		return
	}
	grpcPanicRecoveryHandler := commongrpc.PanicRecoveryGrpc(g.logger)
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
	))
	protoknight.RegisterKnightServiceServer(grpcServer, g)

	reflection.Register(grpcServer)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	go func() {
		healthServer.SetServingStatus(g.serviceName, grpc_health_v1.HealthCheckResponse_SERVING)

		if err := g.registry.Register(g.instanceID, g.serviceName, fmt.Sprintf("%s:%d", g.serviceName, g.port)); err != nil {
			g.logger.Error("Failed to register Service", "error", err.Error())
		}
		if err := grpcServer.Serve(lis); err != nil {
			g.logger.Error("Grpc server failed", "error", err.Error())
		}
	}()

	g.gracefulStop(grpcServer, healthServer)

}
func (g *KnightServer) gracefulStop(grpcServer *grpc.Server, healthServer *health.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	healthServer.SetServingStatus(g.serviceName, grpc_health_v1.HealthCheckResponse_NOT_SERVING)
	g.registry.Deregister(g.instanceID, g.serviceName)
	g.logger.Info("Shutting down server")
	grpcServer.GracefulStop()
	g.logger.Info("Server gracefully stopped")
}

func NewKnightServer(serviceName string, instanceID string, registry discovery.RegistryApi, knightService knight.KnightServiceApi, port int, logger *slog.Logger) *KnightServer {
	return &KnightServer{
		serviceName:   serviceName,
		instanceID:    instanceID,
		knightService: knightService,
		registry:      registry,
		port:          port,
		logger:        logger,
	}
}
