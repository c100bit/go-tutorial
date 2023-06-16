package main

import (
	"log"
	todoapp "todo-app"
	"todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todoapp.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}

}
