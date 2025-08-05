package main

import (
	"net/http"
)

func handleError(w http.ResponseWriter, _ *http.Request) {
	respondWithError(w, "something went wrong", 400)
}
