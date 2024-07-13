package quest

import (
	"context"

	"github.com/google/uuid"
)

type QuestServiceApi interface {
	AddQuest(context.Context, Quest) (*Quest, error)
	GetAssignedQuests(context.Context, uuid.UUID) ([]*Quest, error)
	GetQuest(context.Context, uuid.UUID) (*Quest, error)
	CompleteQuest(context.Context, uuid.UUID) error
}
type QuestRepository interface {
	AddQuest(context.Context, Quest) (*Quest, error)
	GetAssignedQuests(context.Context, uuid.UUID) ([]*Quest, error)
	CompleteQuest(context.Context, uuid.UUID) error
	GetQuest(context.Context, uuid.UUID) (*Quest, error)
}
type KnightService interface {
	GetKnightStatus(uuid.UUID) (bool, error)
}
