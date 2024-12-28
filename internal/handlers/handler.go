package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/patrickmn/go-cache"
)

type Handler struct {
	cache *cache.Cache
}

func NewHandler(cache *cache.Cache) *Handler {
	return &Handler{
		cache: cache,
	}
}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Welcome to the Proxy Server"))
}

var REQUEST_URL = "https://dummyjson.com/products/"

func (h *Handler) GetData(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Cache Status from Cache Middleware" + w.Header().Get("X-Cache") + "\n\n")

	count := r.URL.Query().Get("count")
	if count == "" {
		count = "5"
	}

	resp, err := http.Get(REQUEST_URL + count)
	if err != nil {
		handlerResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"data":    nil,
		})
	}

	defer resp.Body.Close()

	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		handlerResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"data":    nil,
		})
	}

	// set the cache data
	h.cache.Set(count, data, cache.DefaultExpiration)

	handlerResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

func handlerResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
