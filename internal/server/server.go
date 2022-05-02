package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	httpServer http.Server
}

func NewServer() *Server {
	r := mux.NewRouter()
	r.HandleFunc("/hello", sayHello).Methods("GET")

	return &Server{
		http.Server{
			Handler: r,
			Addr: "127.0.0.1:9080",
			ReadTimeout: 15 * time.Second,
			WriteTimeout: 15 * time.Second,
		},
	}
}

func (s *Server) Start() {

	log.Println("Starting server...")

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}

func (s *Server) Stop() {

	log.Println("Stopping server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	s.httpServer.Shutdown(ctx)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello World!",
	})
}
