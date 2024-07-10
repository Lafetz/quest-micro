package quest

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrKntUnavailable = errors.New("knight is currently unavailable")
)

type QuestService struct {
	repo   QuestRepository
	KntSrv KnightService
}

func (srv *QuestService) AddQuest(ctx context.Context, quest Quest) (*Quest, error) {

	canAccept, err := srv.KntSrv.IsAvailable(quest.KnightID)
	if err != nil {
		return nil, err
	}
	if !canAccept {
		return nil, ErrKntUnavailable
	}
	return srv.repo.AddQuest(ctx, quest)
}
func (srv *QuestService) GetAssignedQuests(ctx context.Context, knightId uuid.UUID) ([]*Quest, error) {
	return srv.repo.GetAssignedQuests(ctx, knightId)
}
func (srv *QuestService) GetQuest(ctx context.Context, questId uuid.UUID) (*Quest, error) {
	return srv.repo.GetQuest(ctx, questId)
}
func (srv *QuestService) CompleteQuest(ctx context.Context, questId uuid.UUID) error {
	return srv.repo.CompleteQuest(ctx, questId)
}
