package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}

func (cfg *apiConfig) handleMetrics(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf(
		"<html>\n"+
			"\n<body>\n"+
			"<h1>Welcome, Chirpy Admin</h1>\n"+
			"<p>Chirpy has been visited %v times!</p>\n"+
			"</body>\n\n"+
			"</html>", cfg.fileServerHits)))
	if err != nil {
		return
	}
}

func validateChirp(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	type Chirp struct {
		Body string `json:"body"`
	}

	resp := Chirp{}
	err := decoder.Decode(&resp)
	if err != nil {
		i, _ := json.Marshal(struct {
			Error string `json:"error"`
		}{
			Error: "Something went wrong",
		},
		)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write(i)
		if err != nil {
			return
		}
		return
	}

	if len(resp.Body) > 140 {
		i, _ := json.Marshal(struct {
			Error string `json:"error"`
		}{
			Error: "Chirp is too long",
		},
		)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write(i)
		if err != nil {
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)

	toCheck := strings.Split(resp.Body, " ")

	for i, str := range toCheck {
		if strings.ToLower(str) == "kerfuffle" || strings.ToLower(str) == "sharbert" || strings.ToLower(str) == "fornax" {
			toCheck[i] = "****"
		}
	}

	dat, err := json.Marshal(struct {
		Body string `json:"cleaned_body"`
	}{
		Body: strings.Join(toCheck, " "),
	},
	)
	if err != nil {
		return
	}

	_, err1 := w.Write(dat)
	if err1 != nil {
		return
	}
}
