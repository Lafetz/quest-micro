package quest

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrKntUnavailable = errors.New("knight is currently unavailable")
	ErrNotFound       = errors.New("quest not found")
)

type QuestService struct {
	repo   QuestRepository
	KntSrv KnightService
}

func (srv *QuestService) AddQuest(ctx context.Context, quest Quest) (*Quest, error) {

	canAccept, err := srv.KntSrv.GetKnightStatus(ctx, quest.KntUsername)
	if err != nil {
		return nil, err
	}
	if !canAccept {
		return nil, ErrKntUnavailable
	}
	return srv.repo.AddQuest(ctx, quest)
}
func (srv *QuestService) GetAssignedQuests(ctx context.Context, kntUsername string) ([]*Quest, error) {
	return srv.repo.GetAssignedQuests(ctx, kntUsername)
}
func (srv *QuestService) GetQuest(ctx context.Context, questId uuid.UUID) (*Quest, error) {
	return srv.repo.GetQuest(ctx, questId)
}
func (srv *QuestService) CompleteQuest(ctx context.Context, questId uuid.UUID) error {
	return srv.repo.CompleteQuest(ctx, questId)
}
func NewQuestService(repo QuestRepository, KntSrv KnightService) *QuestService {
	return &QuestService{
		repo:   repo,
		KntSrv: KntSrv,
	}
}
