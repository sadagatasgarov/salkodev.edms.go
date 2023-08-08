package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create index on simple field (non-unique)
func CreateCollectionIndexOnField(ctx context.Context, collection *mongo.Collection, fieldName string) error {

	indexModel := mongo.IndexModel{Keys: bson.D{{Key: fieldName, Value: 1}}}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	return err
}

// Create unique index on one field
func CreateCollectionUniqueIndexOnField(ctx context.Context, collection *mongo.Collection, fieldName string) error {

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: fieldName, Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	return err
}
