package web

import (
	"github.com/google/uuid"
	quest "github.com/lafetz/quest-demo/services/quest/core"
)

type Quest struct {
	ID          uuid.UUID         `json:"id"`
	Owner       string            `json:"owner"`
	KnightID    uuid.UUID         `json:"knightId"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Status      quest.QuestStatus `json:"status"`
}

func toJsonQuest(qst *quest.Quest) Quest {
	return Quest{
		ID:          qst.ID,
		Owner:       qst.Owner,
		KnightID:    qst.KnightID,
		Name:        qst.Name,
		Description: qst.Description,
		Status:      qst.Status,
	}
}
func toJsonQuests(entities []*quest.Quest) []Quest {
	quests := make([]Quest, len(entities))
	for i, entity := range entities {
		quests[i] = toJsonQuest(entity)
	}
	return quests
}
