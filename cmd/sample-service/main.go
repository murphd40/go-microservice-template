package main

import (
	"os"
	"os/signal"
	"strings"

	"github.com/murphd40/go-microservice-template/internal/config"
	"github.com/murphd40/go-microservice-template/internal/dao/repository"
	"github.com/murphd40/go-microservice-template/internal/logging"
	"github.com/murphd40/go-microservice-template/internal/server"
	"github.com/murphd40/go-microservice-template/internal/server/handler"
	"github.com/murphd40/go-microservice-template/internal/service"
)

func main() {

	logging.Configure("INFO")

	envMap := make(map[string]string)
	for _, item := range os.Environ() {
		parts := strings.Split(item, "=")
		envMap[parts[0]] = parts[1]
	}

	serverProperties := config.NewServerProperties()
	config.PopulateConfig(envMap, &serverProperties)

	chatMessageRepository := repository.NewChatMessageRepository()
	chatMessageService := service.NewChatMessageService(chatMessageRepository)
	chatMessageHandler := handler.NewChatMessageHandler(chatMessageService)

	s := server.NewServer(&serverProperties, *chatMessageHandler)

	s.Start()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	logging.Info("Shutting down...")

	s.Stop()

	os.Exit(0)

}
