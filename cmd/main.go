package main

import (
	todo "github.com/zhavzharik/REST-API-in-Go"
	"github.com/zhavzharik/REST-API-in-Go/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
