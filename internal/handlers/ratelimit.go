package handlers

import "net/http"

func (h *Handler) GetDataRatelimited(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Ratelimit-Limit", "100")
	w.Header().Set("X-Ratelimit-Remaining", "99")
	w.Header().Set("X-Ratelimit-Reset", "60")
	HandlerResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    "Ratelimit data",
	})
}
