package main

import (
	"log"
	"net/http"
)

type apiConfig struct {
	fileServerHits int
}

func main() {
	cfg := apiConfig{fileServerHits: 0}
	mux := http.NewServeMux()
	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	mux.HandleFunc("/healthz", handler)
	mux.HandleFunc("/metrics", cfg.handleMetrics)

	corsMux := middlewareCors(mux)

	server := &http.Server{
		Handler: corsMux,
		Addr:    ":8080",
	}

	log.Printf("Serving on port: %s\n", "8080")
	log.Fatal(server.ListenAndServe())

}
