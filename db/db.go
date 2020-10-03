package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goCore/config"

	"log"
)


var CNX *mongo.Client

func InitConnection(config *config.FlagConfig){

	clientOptions := options.Client().ApplyURI(config.DbUrl)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	CNX = client

}

func GetCollection(collection string) *mongo.Collection {
	return CNX.Database(config.Fconfig.DbDatabase).Collection(collection)
}


