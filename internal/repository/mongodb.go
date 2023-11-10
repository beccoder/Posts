package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongoDB(EnvMongoURI string) (*mongo.Client, error) {
	credential := options.Credential{
		Username: "admin",
		Password: "qwerty",
	}
	clientOptions := options.Client().ApplyURI(EnvMongoURI).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
	return client, nil
}

func InitSchemas(client *mongo.Client) error {
	usersColl := client.Database("blogs").Collection("users")
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"username": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	indexView := usersColl.Indexes()
	_, err := indexView.CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}
	return nil
}
