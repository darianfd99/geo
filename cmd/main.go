package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/darianfd99/geo/pkg/handler"
	"github.com/darianfd99/geo/pkg/repository"
	"github.com/darianfd99/geo/pkg/server"
	"github.com/darianfd99/geo/pkg/service"
	"github.com/sirupsen/logrus"
)

var port = os.Getenv("PORT")

func main() {
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := server.NewServer()

	go func() {
		if err := srv.Run(port, handlers.InitRoutes()); err != nil {
			log.Fatalf("error ocurred while running http server: %s", err.Error())
		}
	}()

	log.Println("app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("app shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error ocurred on server shutting down: %s", err.Error())
	}

}
