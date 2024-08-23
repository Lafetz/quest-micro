package client

import (
	commongrpc "github.com/lafetz/quest-micro/common/grpc"
	knightv1 "github.com/lafetz/quest-micro/proto/gen/knight/v1"

	"github.com/sony/gobreaker"
)

func NewKnightClient(addr string, cb *gobreaker.CircuitBreaker) (knightv1.KnightServiceClient, error) {
	conn, err := commongrpc.NewGRPCClient(addr, cb)

	if err != nil {
		return nil, err
	}

	return knightv1.NewKnightServiceClient(conn), nil
}
