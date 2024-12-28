package main

// https://roadmap.sh/projects/caching-server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/karthikeyaspace/proxy-go/internal/handlers"
	"github.com/karthikeyaspace/proxy-go/internal/middleware"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

var (
	PORT = 8080
)

func (s *APIServer) Start() error {

	router := http.NewServeMux()
	handler := handlers.NewHandler()

	router.HandleFunc("GET /", handler.HomeHandler)
	router.HandleFunc("GET /cached-data", handler.GetDataCached)
	router.HandleFunc("GET /ratelimit-data", handler.GetDataRatelimited)

	server := &http.Server{
		Addr:    s.addr,
		Handler: middleware.Logger(router),
	}

	log.Printf("Server started: http://localhost%s", s.addr)
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func main() {
	server := NewAPIServer(fmt.Sprintf(":%d", PORT))
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
