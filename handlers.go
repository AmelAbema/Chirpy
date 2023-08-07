package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}

func (cfg *apiConfig) handleMetrics(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileServerHits)))
	if err != nil {
		return
	}
}
