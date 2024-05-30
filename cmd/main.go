package main

import (
	"log"

	"github.com/maximus969/users-app"	
	"github.com/maximus969/users-app/pkg/repository"
	"github.com/maximus969/users-app/pkg/service"
	"github.com/maximus969/users-app/pkg/handler"
)

func main () {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)


	srv := new(users.Server)	

	if error := srv.Run("8000", handlers.InitRoutes()); error != nil {
		log.Fatalf("error occured while running http server: %s", error.Error())
	}
}