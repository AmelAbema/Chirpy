package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type apiConfig struct {
	fileServerHits int
}

func main() {
	cfg := apiConfig{fileServerHits: 0}

	router := chi.NewRouter()

	router.Handle("/app/*", cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	router.Handle("/app", cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	router.Get("/healthz", handler)
	router.Get("/metrics", cfg.handleMetrics)

	corsMux := middlewareCors(router)

	server := &http.Server{
		Handler: corsMux,
		Addr:    ":8080",
	}

	log.Printf("Serving on port: %s\n", "8080")
	log.Fatal(server.ListenAndServe())

}
