package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Ім'я колекції Users (користувачі системи)
const UsersCollectionName = "Users"

// Підключитися до Mongodb
func _Connect() *mongo.Client {

	uri := os.Getenv("MONGODB_URI_SALKODEV")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Failed to Connect")
		return nil
	}
	log.Println("Successfully Connected to the Mongodb")
	return client
}

var DBClient *mongo.Client = _Connect()

// Отримати колекцію Users бази даних
func Users() *mongo.Collection {
	dbName := os.Getenv("MONGODB_SALKODEV_EDMS")
	if dbName == "" {
		log.Fatal("You must set your 'MONGODB_SALKODEV_EDMS' environmental variable")
		panic("You must set your 'MONGODB_SALKODEV_EDMS' environmental variable")
	}

	collection := DBClient.Database(dbName).Collection(UsersCollectionName)
	return collection
}
