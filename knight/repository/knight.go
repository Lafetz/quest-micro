package repository

import (
	"context"
	"regexp"
	"strings"

	"github.com/google/uuid"
	commonerrors "github.com/lafetz/quest-micro/common/errors"
	knight "github.com/lafetz/quest-micro/knight/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type knightMongo struct {
	Id       uuid.UUID `bson:"_id,omitempty"`
	Name     string    `bson:"name"`
	Email    string    `bson:"email" unique:"true"`
	IsActive bool      `bson:"isActive"`
}

// domain converts knightMongo to Knight
func (u *knightMongo) domain() *knight.Knight {
	return &knight.Knight{
		Id:       u.Id,
		Name:     u.Name,
		Email:    u.Email,
		IsActive: u.IsActive,
	}
}

func newUserMongo(u *knight.Knight) *knightMongo {
	return &knightMongo{
		Id:       u.Id,
		Name:     u.Name,
		Email:    u.Email,
		IsActive: u.IsActive,
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
func (store *Store) GetKnights(ctx context.Context) ([]*knight.Knight, error) {
	var knightsMongo []knightMongo
	cursor, err := store.knights.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var knightData knightMongo
		if err := cursor.Decode(&knightData); err != nil {
			return nil, err
		}
		knightsMongo = append(knightsMongo, knightData)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	var knights []*knight.Knight
	for _, knightMongo := range knightsMongo {
		knights = append(knights, knightMongo.domain())
	}

	return knights, nil
}

func (store *Store) UpdateStatus(ctx context.Context, username string, active bool) error {
	filter := bson.D{{Key: "username", Value: username}}
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

func (store *Store) DeleteKnight(ctx context.Context, username string) error {
	filter := bson.D{{Key: "username", Value: username}}

	result, err := store.knights.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return commonerrors.ErrKnightNotFound
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
