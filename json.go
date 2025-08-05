package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, payload any, code int) {
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
	type ErrorResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w, ErrorResponse{Error: message}, code)
}
