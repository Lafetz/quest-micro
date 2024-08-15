package quest

import (
	"context"

	"github.com/google/uuid"
)

type QuestServiceApi interface {
	AddQuest(context.Context, Quest) (*Quest, error)
	GetAssignedQuests(context.Context, string) ([]*Quest, error)
	GetQuest(context.Context, uuid.UUID) (*Quest, error)
	CompleteQuest(context.Context, uuid.UUID) error
}
type QuestRepository interface {
	AddQuest(context.Context, Quest) (*Quest, error)
	GetAssignedQuests(context.Context, string) ([]*Quest, error)
	CompleteQuest(context.Context, uuid.UUID) error
	GetQuest(context.Context, uuid.UUID) (*Quest, error)
}
type KnightService interface {
	GetKnightStatus(context.Context, string) (bool, error)
}
