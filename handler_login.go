package main

import (
	"encoding/json"
	"github.com/AmelAbema/Chirpy/internal/auth"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	type response struct {
		User
		Token  string `json:"token"`
		RToken string `json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := cfg.DB.GetUserByEmail(params.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get user")
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour, "chirpy-access")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create JWT")
		return
	}
	rToken, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour*24*60, "chirpy-refresh")

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:    user.ID,
			Email: user.Email,
			IsRed: user.IsRed,
		},
		Token:  token,
		RToken: rToken,
	})
}
