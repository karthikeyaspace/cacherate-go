package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

type Handler struct {
	cache *cache.Cache
}

func NewHandler() *Handler {
	return &Handler{
		cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Welcome to the Proxy Server"))
}


// Helper function to send response in JSON format
func HandlerResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}