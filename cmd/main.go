package main

import (
	todo "Go_new_api"
	"Go_new_api/pkg/handler"
	"Go_new_api/pkg/repository"
	"Go_new_api/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
