package main

// https://roadmap.sh/projects/caching-server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/karthikeyaspace/proxy-go/internal/handlers"
	"github.com/karthikeyaspace/proxy-go/internal/middleware"
	"github.com/patrickmn/go-cache"
)

type APIServer struct {
	addr  string
	cache *cache.Cache
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr:  addr,
		cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

var (
	PORT = 8080
)

func (s *APIServer) Start() error {

	router := http.NewServeMux()
	handler := handlers.NewHandler(s.cache)

	cacheMiddleware := middleware.NewCacheMiddleware(s.cache)

	router.HandleFunc("GET /", handler.HomeHandler)
	router.HandleFunc("GET /cached-data", cacheMiddleware.Cache(handler.GetData))
	router.HandleFunc("GET /ratelimit-data", handler.GetData)

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
