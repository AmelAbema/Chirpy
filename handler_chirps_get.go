package main

import (
	"github.com/AmelAbema/Chirpy/internal/database"
	"net/http"
	"sort"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	chirpIDString := chi.URLParam(r, "chirpID")
	chirpID, err := strconv.Atoi(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID")
		return
	}

	dbChirp, err := cfg.DB.GetChirp(chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp")
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:       dbChirp.ID,
		AuthorID: dbChirp.AuthorID,
		Body:     dbChirp.Body,
	})
}

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {

	var chirps []Chirp

	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps")
		return
	}

	s := r.URL.Query().Get("author_id")
	if s != "" {
		id, err := strconv.Atoi(s)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps")
			return
		}
		chirps = retrieveChirps(id, dbChirps, chirps)
	} else {
		chirps = retrieveChirps(0, dbChirps, chirps)
	}

	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].ID < chirps[j].ID
	})

	respondWithJSON(w, http.StatusOK, chirps)
}

func retrieveChirps(id int, dbChirps []database.Chirp, chirps []Chirp) []Chirp {

	for _, dbChirp := range dbChirps {
		if id == 0 {
			chirps = append(chirps, Chirp{
				ID:       dbChirp.ID,
				AuthorID: dbChirp.AuthorID,
				Body:     dbChirp.Body,
			})
		} else {
			if dbChirp.AuthorID == id {
				chirps = append(chirps, Chirp{
					ID:       dbChirp.ID,
					AuthorID: dbChirp.AuthorID,
					Body:     dbChirp.Body,
				})
			}
		}

	}
	return chirps
}
