package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/patrickmn/go-cache"
	"github.com/karthikeyaspace/proxy-go/internal/helpers"
)

var REQUEST_URL = "https://dummyjson.com/products/"

func (h *Handler) GetDataCached(w http.ResponseWriter, r *http.Request) {
	count := r.URL.Query().Get("count")
	if count == "" {
		count = "5"
	}

	if data, ok := h.cache.Get(count); ok {
		w.Header().Set("X-Cache", "HIT")
		fmt.Printf("Cache hit for key: %s\n", count)
		helpers.HandleResponse(w, http.StatusOK, map[string]interface{}{
			"success": true,
			"data":    data,
		})
		return
	}

	// cache status - HIT or MISS
	fmt.Printf("Cache miss for key: %s\n", count)
	w.Header().Set("X-Cache", "MISS")

	resp, err := http.Get(REQUEST_URL + count)
	if err != nil {
		helpers.HandleResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"data":    nil,
		})
		return
	}

	defer resp.Body.Close()

	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		helpers.HandleResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"data":    nil,
		})
		return
	}

	// set the cache data
	h.cache.Set(count, data, cache.DefaultExpiration)

	helpers.HandleResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}
