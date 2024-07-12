package repository

import (
	"context"
	"regexp"
	"strings"

	"github.com/google/uuid"
	commonerrors "github.com/lafetz/quest-demo/common/errors"
	knight "github.com/lafetz/quest-demo/services/knight/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type knightMongo struct {
	Id       uuid.UUID `bson:"_id,omitempty"`
	Username string    `bson:"username" unique:"true"`
	Email    string    `bson:"email" unique:"true"`
	Password []byte    `bson:"password"`
	IsActive bool      `bson:"isActive"`
}

// domain converts knightMongo to Knight
func (u *knightMongo) domain() *knight.Knight {
	return &knight.Knight{
		Id:       u.Id,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		IsActive: u.IsActive,
	}
}

func newUserMongo(u *knight.Knight) *knightMongo {
	return &knightMongo{
		Id:       u.Id,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (store *Store) AddKnight(ctx context.Context, knightData *knight.Knight) (*knight.Knight, error) {
	u := newUserMongo(knightData)
	_, err := store.knights.InsertOne(ctx, u)
	if err != nil {
		if mongoErr, ok := err.(mongo.WriteException); ok {

			for _, writeErr := range mongoErr.WriteErrors {
				if writeErr.Code == 11000 {

					key := extractDuplicateKey(writeErr.Message)
					switch key {
					case "email":
						return nil, knight.ErrEmailUnique
					case "username":
						return nil, knight.ErrUsernameUnique
					default:
						return nil, err

					}
				}
			}
		} else {
			return nil, err
		}
		return knightData, nil
	}
	return knightData, nil
}
func (store *Store) GetKnight(ctx context.Context, username string) (*knight.Knight, error) {
	var knightData knightMongo
	err := store.knights.FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(&knightData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, commonerrors.ErrKnightNotFound
		}
		return nil, err
	}

	return knightData.domain(), nil
}

func (store *Store) UpdateStatus(ctx context.Context, knightID string, active bool) error {
	filter := bson.D{{Key: "_id", Value: knightID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "isActive", Value: active}}}}

	var knightData knightMongo
	err := store.knights.FindOneAndUpdate(ctx, filter, update).Decode(&knightData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return commonerrors.ErrKnightNotFound
		}
		return err
	}
	return nil
}

func extractDuplicateKey(errorMessage string) string {

	pattern := `index: ([a-zA-Z0-9_]+)_\d+ dup key`

	re := regexp.MustCompile(pattern)

	match := re.FindStringSubmatch(errorMessage)
	if len(match) < 2 {
		return ""
	}

	indexName := match[1]

	parts := strings.Split(indexName, "_")
	if len(parts) > 0 {
		fieldName := parts[0]
		return fieldName
	}

	return ""
}
