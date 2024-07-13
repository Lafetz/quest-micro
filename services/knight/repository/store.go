package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	knights *mongo.Collection
}

func NewDb(url string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewStore(client *mongo.Client) (*Store, error) {

	knights := client.Database("knight").Collection("knights")
	err := createUniqueIndex(context.Background(), knights, "email")
	if err != nil {
		return nil, err
	}
	err = createUniqueIndex(context.Background(), knights, "username")
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
