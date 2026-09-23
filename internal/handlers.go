package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ahmadnull/rss-aggregator-go/internal/database"
	"github.com/google/uuid"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

func HandlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}

func (apiCfg *APIConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUsertoUser(user))
}

func (apiCfg *APIConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := getAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUsertoUser(user))
}
