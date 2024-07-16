package repository

import (
	"fmt"

	quest "github.com/lafetz/quest-demo/quest/core"
	"github.com/lafetz/quest-demo/quest/repository/gen"
)

type ErrStatus string

func (e ErrStatus) Error() string {
	return fmt.Sprintf("unknown status: %s", string(e))
}

func toQuestStatus(status gen.QuestStatus) (quest.QuestStatus, error) {
	var questStatusValue quest.QuestStatus

	switch status {
	case "active":
		questStatusValue = quest.QuestStatusActive
	case "completed":
		questStatusValue = quest.QuestStatusCompleted
	case "failed":
		questStatusValue = quest.QuestStatusFailed
	default:
		return "", ErrStatus(status)
	}

	return questStatusValue, nil
}
func mapQuest(q gen.Quest) (*quest.Quest, error) {
	status, err := toQuestStatus(q.Status)
	if err != nil {
		return nil, err
	}

	return &quest.Quest{
		ID:          q.ID,
		Owner:       q.Owner,
		KntUsername: q.KnightUsername,
		Name:        q.Name,
		Description: q.Description,
		Status:      status,
	}, nil
}
func mapQuests(q []gen.Quest) ([]*quest.Quest, error) {
	qsts := []*quest.Quest{}
	for _, qst := range q {
		qstEnt, err := mapQuest(qst)
		if err != nil {
			return nil, err
		}
		qsts = append(qsts, qstEnt)
	}
	return qsts, nil
}
