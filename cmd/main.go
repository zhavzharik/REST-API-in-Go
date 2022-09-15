package main

import (
	todo "github.com/zhavzharik/REST-API-in-Go"
	"github.com/zhavzharik/REST-API-in-Go/pkg/handler"
	"github.com/zhavzharik/REST-API-in-Go/pkg/repository"
	"github.com/zhavzharik/REST-API-in-Go/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
