package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/patrickmn/go-cache"
)

type CacheMiddleware struct {
	cache *cache.Cache
}

func NewCacheMiddleware(cache *cache.Cache) *CacheMiddleware {
	return &CacheMiddleware{
		cache: cache,
	}
}

func (c *CacheMiddleware) Cache(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("count")

		if data, ok := c.cache.Get(key); ok {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Cache", "HIT")
			log.Printf("Cache hit for key: %s\n\n", key)
			json.NewEncoder(w).Encode(data)
			return
		}

		// Cache Miss
		log.Printf("Cache miss for key: %s", key)
		w.Header().Set("X-Cache", "MISS")

		next(w, r)
	}
}
