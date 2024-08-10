package client

import (
	"context"
	"log/slog"
	"time"

	commonerrors "github.com/lafetz/quest-micro/common/errors"
	protoknight "github.com/lafetz/quest-micro/proto/gen"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KnightServiceClient struct {
	logger *slog.Logger
	client protoknight.KnightServiceClient
	cb     *gobreaker.CircuitBreaker
}

func (k *KnightServiceClient) GetKnightStatus(ctx context.Context, email string) (bool, error) {

	req := &protoknight.KnightStatusReq{Email: email}
	res, err := k.cb.Execute(func() (interface{}, error) {
		return k.client.GetKnightStatus(ctx, req)
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {

			return false, commonerrors.ErrKnightNotFound
		}
		return false, err
	}
	return res.(*protoknight.KnightStatusRes).IsActive, nil
}

func newCb() *gobreaker.CircuitBreaker {
	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "Knight.GetKnightStatus",
		Timeout: 60 * time.Second,
	})
}
func NewKntClient(logger *slog.Logger, client protoknight.KnightServiceClient) *KnightServiceClient {
	return &KnightServiceClient{
		logger: logger,
		client: client,
		cb:     newCb(),
	}
}
