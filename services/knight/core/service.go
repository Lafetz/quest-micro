package knight

import (
	"context"
	"errors"
)

var (
	ErrUsernameUnique = errors.New("an account with this username exists")
	ErrDelete         = errors.New("failed to Delete user")
	ErrEmailUnique    = errors.New("an account with this email exists")
)

type KnightService struct {
	repo KnightRepository
}

func (k *KnightService) GetKnight(ctx context.Context, name string) (*Knight, error) {
	return k.repo.GetKnight(ctx, name)
}

func (k *KnightService) AddKnight(ctx context.Context, knight Knight) (*Knight, error) {
	return k.repo.AddKnight(ctx, knight)
}
