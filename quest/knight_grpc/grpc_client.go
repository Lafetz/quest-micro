package client

import (
	commongrpc "github.com/lafetz/quest-micro/common/grpc"
	protoknight "github.com/lafetz/quest-micro/proto/gen"
)

func NewKnightClient(remoteAddr string) (protoknight.KnightServiceClient, error) {
	conn, err := commongrpc.NewGRPCClient(remoteAddr)
	if err != nil {
		return nil, err
	}
	return protoknight.NewKnightServiceClient(conn), nil
}
