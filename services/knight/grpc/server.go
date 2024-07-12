package grpcserver

import (
	"fmt"
	"log"
	"net"
	"strconv"

	protoknight "github.com/lafetz/quest-demo/proto/knight"
	knight "github.com/lafetz/quest-demo/services/knight/core"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	service knight.KnightServiceApi
	port    int
	protoknight.UnimplementedKnightServiceServer
}

func (g *GrpcServer) Run() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(g.port))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	grpcServer := grpc.NewServer()
	protoknight.RegisterKnightServiceServer(grpcServer, g)
	fmt.Println("Starting Server on port ", g.port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Grpc server failed :", err)
	}
}
func NewGrpcServer(service knight.KnightServiceApi, port int) *GrpcServer {
	return &GrpcServer{
		service: service,
		port:    port,
	}
}
