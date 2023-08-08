package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Змінна оточення яка містить URI до MongoDB
const MongoDBURIEnv = "SALKODEV_EDMS_MONGODB_URI"

// Змінна оточення яка містить назву бази в MongoDB
const MongoDBDataBaseEnv = "SALKODEV_EDMS_MONGODB_DATABASE"

// Підключитися до Mongodb
func _Connect() *mongo.Client {

	uri := os.Getenv(MongoDBURIEnv)
	if uri == "" {
		log.Fatal("You must set your '", MongoDBURIEnv, "' environmental variable")
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

func DataBase() *mongo.Database {
	dbName := os.Getenv(MongoDBDataBaseEnv)
	if dbName == "" {
		log.Fatal("You must set your '", MongoDBDataBaseEnv, "' environmental variable")
		panic("You must set your '" + MongoDBDataBaseEnv + "' environmental variable")
	}

	return DBClient.Database(dbName)
}
