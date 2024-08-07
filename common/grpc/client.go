package commongrpc

import (
	"context"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcConnection(host string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	retryOpts := grpc.WithUnaryInterceptor(
		grpc_retry.UnaryClientInterceptor(grpc_retry.WithCodes(codes.Internal),
			grpc_retry.WithMax(5),
			grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second))))
	var Dialopts []grpc.DialOption
	Dialopts = append(Dialopts, retryOpts)
	Dialopts = append(Dialopts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.DialContext(ctx, host, Dialopts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
