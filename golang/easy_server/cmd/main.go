package main

import (
	"/handler"
	"log"
	"todo"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf(err.error)
	}
}
