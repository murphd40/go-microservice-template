package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/murphd40/go-microservice-template/internal/server"
)

func main() {

	s := server.NewServer()

	s.Start()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	s.Stop()

	log.Println("Shutting down...")

	os.Exit(0)

}
