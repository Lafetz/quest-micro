package grpcserver

import (
	"log"
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
		log.Fatalln("Failed to listen:", err)
	}
	grpcPanicRecoveryHandler := commongrpc.PanicRecoveryGrpc(g.logger)
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
	))
	protoknight.RegisterKnightServiceServer(grpcServer, g)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			g.logger.Error("Grpc server failed", "error", err.Error())
		}
	}()
	g.gracefulStop(grpcServer)

}
func (g *GrpcServer) gracefulStop(grpcServer *grpc.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	g.logger.Info("Shutting down server...")
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
