package helpers

import (
	"encoding/json"
	"net/http"
)


func HandleResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}