package server

import (
	"encoding/json"
	"net/http"

	"preparation/go-projects/govault/internal/response"
	"preparation/go-projects/govault/internal/storage"
)

func SetKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
	}

	var req map[string]string
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	key, value := req["key"], req["value"]
	if key == "" || value == "" {
		http.Error(w, "Missing key or value", http.StatusBadRequest)
	}

	storage.Set(key, value)
	w.WriteHeader(http.StatusCreated)
}

func GetKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
	}

	value, ok := storage.Get(key)
	if !ok {
		response.NotFound(w, "Key not found")
	}

	w.Write([]byte(value))
}

func DeleteKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.JSONResponse(w, http.StatusBadRequest, "Invalid request method")
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		response.JSONResponse(w, http.StatusBadRequest, "Missing key")
	}

	ok := storage.Delete(key)
	if !ok {
		response.NotFound(w, "Key not found")
	}

	response.NoContent(w)
}
