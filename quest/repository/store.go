package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/lafetz/quest-demo/quest/repository/gen"
	_ "github.com/lib/pq"
)

type Store struct {
	quests *gen.Queries
}

func NewDb(db *sql.DB) *Store {

	quests := gen.New(db)
	return &Store{
		quests: quests,
	}
}
func OpenDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
