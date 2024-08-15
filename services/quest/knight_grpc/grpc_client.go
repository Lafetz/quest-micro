package client

import (
	commongrpc "github.com/lafetz/quest-micro/common/grpc"
	protoknight "github.com/lafetz/quest-micro/proto/gen"
	"github.com/sony/gobreaker"
)

func NewKnightClient(addr string, cb *gobreaker.CircuitBreaker) (protoknight.KnightServiceClient, error) {
	conn, err := commongrpc.NewGRPCClient(addr, cb)

	if err != nil {
		return nil, err
	}

	return protoknight.NewKnightServiceClient(conn), nil
}
