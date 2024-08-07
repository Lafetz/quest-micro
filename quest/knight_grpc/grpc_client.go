package client

import (
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	protoknight "github.com/lafetz/quest-micro/proto/knight"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient(remoteAddr string) (protoknight.KnightServiceClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(grpc_retry.WithCodes(codes.Internal), grpc_retry.WithMax(5), grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)))))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(remoteAddr, opts...)
	if err != nil {
		return nil, err
	}

	return protoknight.NewKnightServiceClient(conn), nil
}
