package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mbarek-hani/rssagg/internal/auth"
	"github.com/mbarek-hani/rssagg/internal/database"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, "error parsing json", 400)
		return
	}

	user, err := apiCfg.Db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		respondWithError(w, "something went wrong, couldn't create the user", 500)
		return
	}
	respondWithJson(w, databaseUserToUser(user), 201)
}

func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {
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
	respondWithJson(w, user, 200)
}
