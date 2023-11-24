package internal

import (
	"Blogs"
	_ "Blogs/docs"
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

// @title Blogs Server API
// @version 1.0
// @description Blogs Server in Go using Gin framework

// @host localhost:8086
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func Run() {
	if err := LoadEnvConfig(); err != nil {
		log.Fatal(err)
	}

	dbURI := viper.GetString("MONGO.PROTOCOL") + "://" + viper.GetString("MONGO.HOST") + ":" + viper.GetString("MONGO.PORT")
	client, err := repository.ConnectMongoDB(dbURI)
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
		httpAddr := viper.GetString("HTTP.HOST") + ":" + viper.GetString("HTTP.PORT")
		if err := server.Run(httpAddr, handlers.InitRoutes()); err != nil {
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

func InitConfig() error {
	viper.AddConfigPath("configs")
	if os.Getenv("RUN_MODE") == "test" {
		viper.SetConfigName("test_config")
	} else {
		viper.SetConfigName("prod_config")
	}
	return viper.ReadInConfig()
}

func LoadEnvConfig() error {
	err := InitConfig()
	if err != nil {
		return err
	}
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}

func HandleCLA(args []string) {
	if len(os.Args) != 4 {
		log.Fatalf("invalid args format:\norder of arguments to create superadmin: username password role")
	}

	username := args[1]
	password := args[2]
	role := args[3]

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	if err := InitConfig(); err != nil {
		log.Fatal(err)
	}

	dbURI := viper.GetString("MONGO.PROTOCOL") + "://" + viper.GetString("MONGO.HOST") + ":" + viper.GetString("MONGO.PORT")
	client, err := repository.ConnectMongoDB(dbURI)
	err = repository.InitSchemas(client)
	if err != nil {
		panic(err)
	}
	repos := repository.NewRepository(client)
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
