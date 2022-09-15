package main

import (
	"github.com/spf13/viper"
	todo "github.com/zhavzharik/REST-API-in-Go"
	"github.com/zhavzharik/REST-API-in-Go/pkg/handler"
	"github.com/zhavzharik/REST-API-in-Go/pkg/repository"
	"github.com/zhavzharik/REST-API-in-Go/pkg/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
