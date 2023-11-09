package main

import (
	"Blogs/internal/repository"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	client, err := repository.ConnectMongoDB(viper.GetString("MONGO_URI"))
	if err != nil {
		panic(err)
	}

	fmt.Println(client)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
