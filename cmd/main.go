package main

import (
	gocrud "go-crud"
	"go-crud/pkg/handler"
	"go-crud/pkg/repository"
	"go-crud/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(gocrud.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("Error starting server:", err.Error())
	}
}
