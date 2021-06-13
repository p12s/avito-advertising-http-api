package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if err := initConfig(); err != nil {
		logger.Fatal("Error init config: " + err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logger.Fatal("Error loading .env variables: " + err.Error())
	}

	// инит пустой бд

	// инит репо
	// сервиса
	// хендлера

	// запуск сервера
	logger.Info("Service is on")

	// shutdown
	logger.Info("Service is off")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
