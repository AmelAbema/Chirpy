package main

import (
	"encoding/json"
	"github.com/AmelAbema/Chirpy/internal/auth"
	"net/http"
)

func (cfg *apiConfig) handlerPolkaWebhooks(w http.ResponseWriter, r *http.Request) {
	token, err1 := auth.GetBearerToken(r.Header)
	if err1 != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find ApiKey")
		return
	}

	if token != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT")
		return
	}

	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID int `json:"user_id"`
		} `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	if params.Event == "user.upgraded" {
		user, err := cfg.DB.GetUser(params.Data.UserID)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Couldn't find user parameters")
			return
		}
		_, err1 := cfg.DB.UpdateUser(user.ID, user.Email, user.HashedPassword, true)
		if err1 != nil {
			return
		}
		respondWithJSON(w, http.StatusOK, User{})
		return
	}

	respondWithJSON(w, http.StatusOK, User{})
}
