package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(EnvMongoURI string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(EnvMongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")
	//collection := client.Database("go_rest_api").Collection("books")
	return client, nil
}
