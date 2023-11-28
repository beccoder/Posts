package internal

import (
	"Blogs"
	_ "Blogs/docs"
	"Blogs/internal/handler"
	"Blogs/internal/repository"
	"Blogs/internal/service"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title Blogs Server API
// @version 1.0
// @description Blogs Server in Go using Gin framework

// @host http://localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func Run(cfg *Blogs.Config) {
	cfg.MakeMongoDBURL()
	client, err := repository.ConnectMongoDB(cfg.MONGODB.URI, cfg.MONGODB.Username, cfg.MONGODB.Password)
	if err != nil {
		panic(err)
	}
	err = repository.InitSchemas(client, cfg.MONGODB.Database)
	if err != nil {
		panic(err)
	}
	repos := repository.NewRepository(client, cfg.MONGODB.Database)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(Blogs.Server)
	go func() {
		cfg.MakeHttpURL()
		if err := server.Run(cfg.Http.URL, handlers.InitRoutes()); err != nil {
			log.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	log.Println("Blogs Started")

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

func HandleCLA(args []string, cfg *Blogs.Config) {
	if len(os.Args) != 4 {
		log.Fatalf("invalid args format:\norder of arguments to create superadmin: username password role")
	}

	username := args[1]
	password := args[2]
	role := args[3]

	cfg.MakeMongoDBURL()
	client, err := repository.ConnectMongoDB(cfg.MONGODB.URI, cfg.MONGODB.Username, cfg.MONGODB.Password)
	err = repository.InitSchemas(client, cfg.MONGODB.Database)
	if err != nil {
		panic(err)
	}
	repos := repository.NewRepository(client, cfg.MONGODB.Database)
	services := service.NewService(repos)
	id, err := services.CreateUser(Blogs.UserModel{
		Role:     role,
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(id)
	}
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error occured on db connection close: %s", err.Error())
	}
}
