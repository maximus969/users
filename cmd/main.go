package main

import (
	"log"

	"github.com/maximus969/users-app"
	"github.com/maximus969/users-app/pkg/handler"
)

func main () {
	srv := new(users.Server)
	handlers := new(handler.Handler)

	if error := srv.Run("8000", handlers.InitRoutes()); error != nil {
		log.Fatalf("error occured while running http server: %s", error.Error())
	}
}