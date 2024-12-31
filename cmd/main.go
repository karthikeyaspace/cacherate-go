package main

// https://roadmap.sh/projects/caching-server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/karthikeyaspace/proxy-go/internal/handlers"
	"github.com/karthikeyaspace/proxy-go/internal/middleware"
	"golang.org/x/time/rate"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

var PORT = 8080

func (s *APIServer) Start() error {

	router := http.NewServeMux()
	handler := handlers.NewHandler()

	ratelimiter := middleware.NewRateLimiter(rate.Every(10*time.Second), 5) // 5 request per 10 seconds

	router.HandleFunc("GET /", handler.HomeHandler)
	router.HandleFunc("GET /caching", handler.GetDataCached)
	router.HandleFunc("GET /ratelimiting", ratelimiter.Ratelimit(http.HandlerFunc(handler.GetDataRatelimited)))

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
