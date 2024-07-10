// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package gen

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
)

type QuestStatus string

const (
	QuestStatusActive    QuestStatus = "active"
	QuestStatusCompleted QuestStatus = "completed"
	QuestStatusFailed    QuestStatus = "failed"
)

func (e *QuestStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = QuestStatus(s)
	case string:
		*e = QuestStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for QuestStatus: %T", src)
	}
	return nil
}

type NullQuestStatus struct {
	QuestStatus QuestStatus
	Valid       bool // Valid is true if QuestStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullQuestStatus) Scan(value interface{}) error {
	if value == nil {
		ns.QuestStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.QuestStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullQuestStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.QuestStatus), nil
}

type Quest struct {
	ID          uuid.UUID
	Owner       string
	KnightID    uuid.UUID
	Name        string
	Description string
	Status      QuestStatus
}
