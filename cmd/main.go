package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/p12s/avito-advertising-http-api"
	"github.com/p12s/avito-advertising-http-api/pkg/handler"
	"github.com/p12s/avito-advertising-http-api/pkg/repository"
	"github.com/p12s/avito-advertising-http-api/pkg/service"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
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
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("DB_USER"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logger.Fatal("Failed to initialize db: " + err.Error())
	}
	// инит репо
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// сервиса
	srv := new(common.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logger.Fatal("error while running http server: " + err.Error())
		}
		logger.Info("Service is on, port " + viper.GetString("port"))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("Service is off")
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Fatal("error occured on server shutting down: " + err.Error())
	}

	if err := db.Close(); err != nil {
		logger.Fatal("error on db connection close: " + err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
