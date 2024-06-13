package main

import (
	"github.com/spf13/viper"
	todo "github.com/zhashcevych/todo-app"
	"github.com/zhashcevych/todo-app/pkg/handlers"
	"github.com/zhashcevych/todo-app/pkg/repository"
	"github.com/zhashcevych/todo-app/pkg/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("erromir occured while running: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
