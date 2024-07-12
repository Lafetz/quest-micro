package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	quest "github.com/lafetz/quest-demo/services/quest/core"
	"github.com/lafetz/quest-demo/services/quest/repository/gen"
)

func (store *Store) AddQuest(ctx context.Context, qst quest.Quest) (*quest.Quest, error) {

	q, err := store.quests.AddQuest(ctx, gen.AddQuestParams{ID: qst.ID, Owner: qst.Owner, KnightID: qst.KnightID,
		Name: qst.Name, Description: qst.Description, Status: gen.QuestStatus(qst.Status)})
	if err != nil {
		return nil, err
	}
	questEnt, err := mapQuest(q)
	return questEnt, err
}

func (store *Store) GetAssignedQuests(ctx context.Context, kid uuid.UUID) ([]*quest.Quest, error) {
	q, err := store.quests.GetAssignedQuests(ctx, kid)
	if err != nil {
		return nil, err
	}
	qsts, err := mapQuests(q)
	if err != nil {
		return nil, err
	}
	return qsts, nil
}

func (store *Store) CompleteQuest(ctx context.Context, questId uuid.UUID) error {
	err := store.quests.CompleteQuest(ctx, questId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return quest.ErrNotFound
		}
		return err
	}
	return nil
}

func (store *Store) GetQuest(ctx context.Context, questId uuid.UUID) (*quest.Quest, error) {
	q, err := store.quests.GetQuest(ctx, questId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, quest.ErrNotFound
		}
		return nil, err
	}
	questEnt, err := mapQuest(q)
	if err != nil {
		return nil, err
	}
	return questEnt, nil
}
