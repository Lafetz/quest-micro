package grpcserver

import (
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	commongrpc "github.com/lafetz/quest-micro/common/grpc"
	knight "github.com/lafetz/quest-micro/knight/core"
	protoknight "github.com/lafetz/quest-micro/proto/knight"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type GrpcServer struct {
	name          string //service name
	version       string //service version
	knightService knight.KnightServiceApi
	port          int
	protoknight.UnimplementedKnightServiceServer
	logger *slog.Logger
}

func (g *GrpcServer) Run() {

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
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	go func() {
		healthServer.SetServingStatus(g.name, grpc_health_v1.HealthCheckResponse_SERVING)
		if err := grpcServer.Serve(lis); err != nil {
			g.logger.Error("Grpc server failed", "error", err.Error())
		}
	}()

	g.gracefulStop(grpcServer, healthServer)

}
func (g *GrpcServer) gracefulStop(grpcServer *grpc.Server, healthServer *health.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	healthServer.SetServingStatus(g.name, grpc_health_v1.HealthCheckResponse_NOT_SERVING)

	g.logger.Info("Shutting down server")
	grpcServer.GracefulStop()
	g.logger.Info("Server gracefully stopped")
}

func NewGrpcServer(knightService knight.KnightServiceApi, port int, logger *slog.Logger) *GrpcServer {
	return &GrpcServer{
		name:          "knight",
		version:       "1.0.0",
		knightService: knightService,
		port:          port,
		logger:        logger,
	}
}
