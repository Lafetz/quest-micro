package quest

import (
	"github.com/google/uuid"
)

type QuestStatus string

const (
	QuestStatusActive QuestStatus = "active"

	QuestStatusCompleted QuestStatus = "completed"
	QuestStatusFailed    QuestStatus = "failed"
)

type Quest struct {
	ID          uuid.UUID
	Owner       string
	KnightID    uuid.UUID
	Name        string
	Description string
	Status      QuestStatus
}

func NewQuest(id uuid.UUID, owner string, knightId uuid.UUID, name string, description string) *Quest {

	if id == uuid.Nil {
		id = uuid.New()
	}
	return &Quest{
		ID:          id,
		Owner:       owner,
		KnightID:    knightId,
		Name:        name,
		Description: description,
		Status:      QuestStatusActive,
	}
}
