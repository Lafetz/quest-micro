package knight

import (
	"github.com/google/uuid"
)

type Knight struct {
	Id       uuid.UUID
	Name     string
	Email    string
	IsActive bool
}

func NewKnight(Name string, email string) *Knight {
	user := &Knight{
		Id:       uuid.New(),
		Name:     Name,
		Email:    email,
		IsActive: true,
	}

	return user
}
