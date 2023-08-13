package main

import (
	"github.com/AmelAbema/Chirpy/internal/auth"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (cfg *apiConfig) handlerChirpsDelete(w http.ResponseWriter, r *http.Request) {

	chirpIDString := chi.URLParam(r, "chirpID")
	chirpID, err := strconv.Atoi(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID")
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT")
		return
	}
	subject, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT")
		return
	}
	userID, err := strconv.Atoi(subject)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't parse user ID")
		return
	}
	chirp, err2 := cfg.DB.GetChirp(chirpID)
	if err2 != nil {
		respondWithError(w, http.StatusForbidden, "Couldn't find chirp")
		return
	}

	if chirp.AuthorID != userID {
		respondWithError(w, http.StatusForbidden, "Couldn't parse user ID")
		return
	}

	err1 := cfg.DB.DeleteChirp(chirpID)
	if err1 != nil {
		respondWithError(w, http.StatusForbidden, "Couldn't delete chirp")
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{})
}
