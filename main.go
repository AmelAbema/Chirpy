package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type apiConfig struct {
	fileServerHits int
	DB             *DB
	JwtToken       string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	jwtSecret := os.Getenv("JWT_SECRET")

	db, err := NewDB("database.json")
	if err != nil {
		log.Fatal(err)
	}
	cfg := apiConfig{
		fileServerHits: 0,
		DB:             db,
		JwtToken:       jwtSecret,
	}

	router := chi.NewRouter()

	router.Handle("/app/*", cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	router.Handle("/app", cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))

	apiRouter := chi.NewRouter()
	router.Mount("/api", apiRouter)
	apiRouter.Get("/healthz", handler)
	apiRouter.Post("/chirps", cfg.handlerChirpsCreate)
	apiRouter.Get("/chirps", cfg.handlerChirpsRetrieve)
	apiRouter.Get("/chirps/{chirpID}", cfg.handleChirpId)
	apiRouter.Post("/users", cfg.handlerUsersCreate)
	apiRouter.Post("/login", cfg.handlerLogin)

	metricsRouter := chi.NewRouter()
	router.Mount("/admin", metricsRouter)
	metricsRouter.Get("/metrics", cfg.handleMetrics)

	corsMux := middlewareCors(router)

	server := &http.Server{
		Handler: corsMux,
		Addr:    ":8080",
	}

	log.Printf("Serving on port: %s\n", "8080")
	log.Fatal(server.ListenAndServe())

}
