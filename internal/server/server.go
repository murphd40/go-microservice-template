package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/murphd40/go-microservice-template/internal/config"
	"github.com/murphd40/go-microservice-template/internal/logging"
	"github.com/murphd40/go-microservice-template/internal/server/handler"
)

type Server struct {
	httpServer http.Server
}

func NewServer(properties *config.ServerProperties, chatMessageHandler handler.ChatMessageHandler) *Server {
	r := mux.NewRouter()
	r.HandleFunc("/hello", sayHello).Methods("GET")
	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/chatmessage", chatMessageHandler.CreateChatMessage).Methods("POST")
	v1.HandleFunc("/chatmessage/{chatMessageId}", chatMessageHandler.GetChatMessageById).Methods("GET")

	return &Server{
		httpServer: http.Server{
			Handler:      r,
			Addr:         fmt.Sprint(":", properties.Port),
			ReadTimeout:  properties.ReadTimeout,
			WriteTimeout: properties.WriteTimeout,
		},
	}
}

func (s *Server) Start() {

	logging.Info("Starting server...")

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			logging.Error(err)
		}
	}()

	logging.Info("Server is listening on ", s.httpServer.Addr)
}

func (s *Server) Stop() {

	logging.Info("Stopping server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	s.httpServer.Shutdown(ctx)
}

func sayHello(w http.ResponseWriter, r *http.Request) {

	data := map[string]any{
		"message": "Hello World!",
		"time":    time.Now(),
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")
	encoder.Encode(data)
}
