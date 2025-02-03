package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"preparation/go-projects/govault/internal/response"
	"preparation/go-projects/govault/internal/storage"
)

var db = storage.NewMultiDB()

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

	db.ActiveDB().Set(key, value)
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

	value, ok := db.ActiveDB().Get(key)
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

	ok := db.ActiveDB().Delete(key)
	if !ok {
		response.NotFound(w, "Key not found")
	}

	response.NoContent(w)
}

func SetKeyTTL(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)

	key, value := data["key"].(string), data["value"].(string)
	ttl := time.Duration(data["ttl"].(int)) * time.Second

	db.ActiveDB().SetWithTTL(key, value, ttl)
	slog.Info("Key set with TTL", "key", key, "ttl", ttl)
	w.WriteHeader(http.StatusCreated)
}

func SelectDB(w http.ResponseWriter, r *http.Request) {
	var data map[string]int
	json.NewDecoder(r.Body).Decode(&data)

	dbNum := data["db"]
	db.SelectDB(dbNum)
	slog.Info("Database switched", "db", dbNum)
	w.WriteHeader(http.StatusOK)
}
