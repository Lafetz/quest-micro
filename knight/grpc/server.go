package grpcserver

import (
	"log"
	"log/slog"
	"net"
	"runtime/debug"
	"strconv"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	knight "github.com/lafetz/quest-demo/knight/core"
	protoknight "github.com/lafetz/quest-demo/proto/knight"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	service knight.KnightServiceApi
	port    int
	protoknight.UnimplementedKnightServiceServer
	logger *slog.Logger
}

func (g *GrpcServer) Run() {

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(g.port))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	//recover
	grpcPanicRecoveryHandler := func(p any) (err error) {
		g.logger.Error("msg", "recovered from panic", "panic", p, "stack", debug.Stack())
		return status.Errorf(codes.Internal, "%s", p)
	}

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
	))
	protoknight.RegisterKnightServiceServer(grpcServer, g)
	g.logger.Info("Starting Server on port " + strconv.Itoa(g.port))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Grpc server failed :", err)
	}
}
func NewGrpcServer(service knight.KnightServiceApi, port int, logger *slog.Logger) *GrpcServer {
	return &GrpcServer{
		service: service,
		port:    port,
		logger:  logger,
	}
}
