package main

import (
	"os"
	"os/signal"

	"github.com/murphd40/go-microservice-template/internal/server"
	"github.com/murphd40/go-microservice-template/internal/server/handler"
	"github.com/murphd40/go-microservice-template/internal/service"
	"github.com/murphd40/go-microservice-template/internal/dao/repository"
	"github.com/murphd40/go-microservice-template/internal/logging"
)

func main() {

	logging.Configure("INFO")

	chatMessageRepository := repository.NewChatMessageRepository()
	chatMessageService := service.NewChatMessageService(chatMessageRepository)
	chatMessageHandler := handler.NewChatMessageHandler(chatMessageService)

	s := server.NewServer(*chatMessageHandler)

	s.Start()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	logging.Info("Shutting down...")

	s.Stop()

	os.Exit(0)

}
