package main

import (
	"fmt"
	"net/http"

	"github.com/mbarek-hani/rssagg/internal/auth"
	"github.com/mbarek-hani/rssagg/internal/database"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) middlwareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, fmt.Sprintf("auth error: %s", err.Error()), 403)
			return
		}
		user, err := apiCfg.Db.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, "couldn't get user", 404)
			return
		}
		handler(w, r, user)
	}
}
