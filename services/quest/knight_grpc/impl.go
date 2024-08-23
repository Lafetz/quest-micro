package client

import (
	"context"
	"log/slog"

	commonerrors "github.com/lafetz/quest-micro/common/errors"
	knightv1 "github.com/lafetz/quest-micro/proto/gen/knight/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KnightServiceClient struct {
	logger *slog.Logger
	client knightv1.KnightServiceClient
}

func (k *KnightServiceClient) GetKnightStatus(ctx context.Context, email string) (bool, error) {

	req := &knightv1.GetKnightStatusRequest{Email: email}
	res, err := k.client.GetKnightStatus(ctx, req)
	if err != nil {
		if status.Code(err) == codes.NotFound {

			return false, commonerrors.ErrKnightNotFound
		}
		return false, err

	}
	return res.IsActive, nil
}

func NewKntClient(logger *slog.Logger, client knightv1.KnightServiceClient) *KnightServiceClient {
	return &KnightServiceClient{
		logger: logger,
		client: client,
	}
}
