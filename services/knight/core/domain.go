package knight

import (
	"time"

	"github.com/google/uuid"
)

type Knight struct {
	Id        uuid.UUID
	Username  string
	Email     string
	Password  []byte
	IsActive  bool
	CreatedAt time.Time
}

func NewUser(username string, email string, password []byte) *Knight {
	user := &Knight{
		Id:        uuid.New(),
		Username:  username,
		Email:     email,
		Password:  password,
		IsActive:  false,
		CreatedAt: time.Now(),
	}

	return user
}
