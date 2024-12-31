package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/karthikeyaspace/proxy-go/internal/helpers"
)

func (h *Handler) GetDataRatelimited(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get(REQUEST_URL + "5")
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

	helpers.HandleResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    data,
	})

}
