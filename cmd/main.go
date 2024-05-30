package main

import (
	"log"

	"github.com/maximus969/users-app"
	"github.com/maximus969/users-app/pkg/handler"
	"github.com/maximus969/users-app/pkg/repository"
	"github.com/maximus969/users-app/pkg/service"
	"github.com/spf13/viper"
)

func main () {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)


	srv := new(users.Server)	

	if error := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); error != nil {
		log.Fatalf("error occured while running http server: %s", error.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}