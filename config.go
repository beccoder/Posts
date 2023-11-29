package Blogs

import (
	"embed"
	"fmt"
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type AppMode string

const (
	TESTING    AppMode = "TESTING"
	PRODUCTION AppMode = "PRODUCTION"
)

//go:embed configs
var configs embed.FS

type Config struct {
	Mode string `env:"APPLICATION_MODE" envDefault:"development"`

	Project struct {
		Name           string `env:"PROJECT_NAME" yaml:"name"`
		Version        string `env:"PROJECT_VERSION" yaml:"version"`
		SwaggerEnabled bool   `env:"PROJECT_SWAGGER_ENABLED" yaml:"swagger_enabled"`
	} `yaml:"project"`

	Http struct {
		Host string `env:"HTTP_HOST" yaml:"host"`
		Port int    `env:"HTTP_PORT" yaml:"port"`

		URL string `env:"HTTP_URL" yaml:"url"`
	} `yaml:"http"`

	MONGODB struct {
		Protocol string `env:"MONGODB_PROTOCOL" envDefault:"mongodb"`
		Host     string `env:"MONGODB_HOST" envDefault:"db"`
		Port     int    `env:"MONGODB_PORT" envDefault:"27017"`
		Username string `env:"MONGODB_USERNAME" envDefault:"admin"`
		Password string `env:"MONGODB_PASSWORD" envDefault:"qwerty"`
		Database string `env:"MONGODB_DATABASE" yaml:"database"`

		URI string `env:"MONGODB_URI"`
	} `yaml:"mongo"`
}

func Load() *Config {
	var cfg Config
	err := godotenv.Load(".env")
	if err != nil && !os.IsNotExist(err) {
		log.Println("failed to load .env file")
	}

	configPath := getConfigPath(AppMode(getAppMode()))

	file, err := configs.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}

	if err := env.Parse(&cfg); err != nil {
		log.Println(err.Error())
		panic("unmarshal from environment error")
	}

	return &cfg
}

func getAppMode() AppMode {
	mode := AppMode(os.Getenv("APPLICATION_MODE"))

	if mode != TESTING {
		mode = PRODUCTION
	}

	return mode
}

func (c *Config) MakeHttpURL() {
	c.Http.URL = fmt.Sprintf("%s:%d", c.Http.Host, c.Http.Port)
}

func (c *Config) MakeMongoDBURL() {
	c.MONGODB.URI = fmt.Sprintf("%s://%s:%d", c.MONGODB.Protocol, c.MONGODB.Host, c.MONGODB.Port)
}

func getConfigPath(appMode AppMode) string {
	suffix := "test_config"
	if appMode == PRODUCTION {
		suffix = "prod_config"
	}

	return fmt.Sprintf("configs/%s.yaml", suffix)
}
