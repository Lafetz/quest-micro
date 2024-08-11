package repository

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	knights *mongo.Collection
}

func NewDb(url string, logger *slog.Logger) (*mongo.Client, func(), error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	return client, func() {

		err := client.Disconnect(context.Background())
		if err != nil {
			logger.Error(err.Error())
		}

	}, nil
}

func NewStore(client *mongo.Client) (*Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	knights := client.Database("knight").Collection("knights")
	err := createUniqueIndex(ctx, knights, "email")
	if err != nil {
		return nil, err
	}

	return &Store{
		knights: knights,
	}, nil
}
func createUniqueIndex(ctx context.Context, collection *mongo.Collection, fieldName string) error {

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: fieldName, Value: 1}},
		Options: options.Index().SetUnique(true),
	} //

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	return nil
}
