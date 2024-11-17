package main

import (
	gocrud "go-crud"
	"go-crud/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(gocrud.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("Error starting server:", err.Error())
	}
}
