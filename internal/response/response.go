package response

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, status int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": message,
	})
}

func NotFound(w http.ResponseWriter, message string) {
	JSONResponse(w, http.StatusNotFound, map[string]string{"error": message})
}

func Success(w http.ResponseWriter, message string) {
	JSONResponse(w, http.StatusOK, map[string]string{"success": message})
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
