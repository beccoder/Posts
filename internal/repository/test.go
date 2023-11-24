package repository

import (
	"Blogs"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func StartTest() *mongo.Client {
	if err := Blogs.LoadEnvConfig(); err != nil {
		log.Fatal(err)
	}

	dbURI := viper.GetString("MONGO.PROTOCOL") + "://" + viper.GetString("MONGO.HOST") + ":" + viper.GetString("MONGO.PORT")
	client, err := ConnectMongoDB(dbURI)
	if err != nil {
		panic(err)
	}
	err = InitSchemas(client)
	if err != nil {
		panic(err)
	}
	return client
}
