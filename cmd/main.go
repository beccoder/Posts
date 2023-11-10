package main

import (
	"Blogs"
	"Blogs/internal/handler"
	"Blogs/internal/repository"
	"Blogs/internal/service"
	"context"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	err = repository.InitSchemas(client)
	if err != nil {
		panic(err)
	}
	repos := repository.NewRepository(client)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(Blogs.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	log.Print("Blogs Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Blogs is Shutting Down")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error occured on server shutting down: %s", err.Error())
	}

	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
