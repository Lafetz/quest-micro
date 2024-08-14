package client

import (
	commongrpc "github.com/lafetz/quest-micro/common/grpc"
	protoknight "github.com/lafetz/quest-micro/proto/gen"
	"github.com/sony/gobreaker"
)

func NewKnightClient(remoteAddr string, cb *gobreaker.CircuitBreaker) (protoknight.KnightServiceClient, error) {
	conn, err := commongrpc.NewGRPCClient(remoteAddr, cb)
	if err != nil {
		return nil, err
	}
	return protoknight.NewKnightServiceClient(conn), nil
}
