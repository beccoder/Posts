package Blogs

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

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
	if err := os.Chdir("/home/beccoder/posts"); err != nil {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		return err
	}

	err := InitConfig()
	if err != nil {
		return err
	}
	return nil
}
