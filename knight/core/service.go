package knight

import (
	"context"
	"errors"
)

var (
	// ErrUsernameUnique = errors.New("an account with this username exists")
	ErrEmailUnique = errors.New("an account with this email exists")
)

type KnightService struct {
	repo KnightRepository
}

func (k *KnightService) KnightStatus(ctx context.Context, name string) (bool, error) {
	knight, err := k.repo.GetKnight(ctx, name)
	if err != nil {
		return false, err
	}
	return knight.IsActive, nil
}
func (k *KnightService) UpdateStatus(ctx context.Context, username string, active bool) error {
	return k.repo.UpdateStatus(ctx, username, active)
}
func (k *KnightService) AddKnight(ctx context.Context, knight *Knight) (*Knight, error) {
	return k.repo.AddKnight(ctx, knight)
}
func (k *KnightService) GetKnights(ctx context.Context) ([]*Knight, error) {
	return k.repo.GetKnights(ctx)
}
func (k *KnightService) DeleteKnight(ctx context.Context, username string) error {
	return k.repo.DeleteKnight(ctx, username)
}
func NewKnightService(repo KnightRepository) *KnightService {
	return &KnightService{repo}
}
