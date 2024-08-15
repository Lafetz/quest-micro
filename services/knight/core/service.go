package knight

import (
	"context"
	"errors"
)

var (
	ErrEmailUnique = errors.New("an account with this email exists")
)

type KnightService struct {
	repo KnightRepository
}

func (k *KnightService) KnightStatus(ctx context.Context, email string) (bool, error) {
	knight, err := k.repo.GetKnight(ctx, email)
	if err != nil {
		return false, err
	}
	return knight.IsActive, nil
}
func (k *KnightService) UpdateStatus(ctx context.Context, email string, active bool) error {
	return k.repo.UpdateStatus(ctx, email, active)
}
func (k *KnightService) AddKnight(ctx context.Context, knight *Knight) (*Knight, error) {
	return k.repo.AddKnight(ctx, knight)
}
func (k *KnightService) GetKnights(ctx context.Context) ([]*Knight, error) {
	return k.repo.GetKnights(ctx)
}
func (k *KnightService) GetKnight(ctx context.Context, email string) (*Knight, error) {
	return k.repo.GetKnight(ctx, email)
}
func (k *KnightService) DeleteKnight(ctx context.Context, email string) error {
	return k.repo.DeleteKnight(ctx, email)
}
func NewKnightService(repo KnightRepository) *KnightService {
	return &KnightService{repo}
}
