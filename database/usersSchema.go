package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Validate Users collection in MongoDB, indexes and others
func ValidateUsersCollection() {

	users := Users()

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := users.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}

	log.Println("UsersCollection schema validated")
}
