package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/smirnoffkin/todo-app/internal/config"
	"github.com/smirnoffkin/todo-app/internal/handler"
	"github.com/smirnoffkin/todo-app/internal/repository"
	"github.com/smirnoffkin/todo-app/internal/service"
	"github.com/smirnoffkin/todo-app/pkg/server"
)

func init() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}
	config.Settings = *config.NewConfig()
}

func main() {
	db, err := repository.GetDBSession(config.Settings.PostgresURL)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(config.Settings.AppPort, handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}
