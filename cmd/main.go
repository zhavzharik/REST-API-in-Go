package main

import (
	todo "github.com/zhavzharik/REST-API-in-Go"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
