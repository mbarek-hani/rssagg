package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		slog.Error("json.Marshal: " + err.Error())
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, message string, code int) {
	respondWithJson(w, code, struct {
		Error string `json:"error"`
	}{
		Error: message,
	})
}
