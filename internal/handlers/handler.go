package handlers

import (
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
