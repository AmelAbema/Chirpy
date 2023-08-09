package main

import "fmt"

type apiConfig struct {
	fileServerHits int
}

func main() {
	//cfg := apiConfig{fileServerHits: 0}
	//
	//router := chi.NewRouter()
	//
	//router.Handle("/app/*", cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	//router.Handle("/app", cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	//
	//apiRouter := chi.NewRouter()
	//router.Mount("/api", apiRouter)
	//apiRouter.Get("/healthz", handler)
	//apiRouter.Post("/chirps", validateChirp)
	//
	//metricsRouter := chi.NewRouter()
	//router.Mount("/admin", metricsRouter)
	//metricsRouter.Get("/metrics", cfg.handleMetrics)
	//
	//corsMux := middlewareCors(router)
	//
	//server := &http.Server{
	//	Handler: corsMux,
	//	Addr:    ":8080",
	//}
	//
	//log.Printf("Serving on port: %s\n", "8080")
	//log.Fatal(server.ListenAndServe())
	_, err := NewDB("C:\\Programming\\MyProjects\\Chirpy\\database.json")
	if err != nil {
		fmt.Printf("error123")
		return
	}
}
