package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Pempho-Mackson-Kapulula/chirpy/internal/api"
	"github.com/Pempho-Mackson-Kapulula/chirpy/internal/config"
	"github.com/Pempho-Mackson-Kapulula/chirpy/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	cfg := config.Load()

	dbConn, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close() // Ensure pool closes cleanly on exit

	if err := dbConn.Ping(); err != nil {
		log.Fatal(err)
	}

	apiCfg := api.Config{
		DB:       database.New(dbConn),
		Platform: cfg.Platform,
	}

	mux := http.NewServeMux()

	// Static asset delivery
	fs := http.FileServer(http.Dir("."))
	mux.Handle("/app/", http.StripPrefix("/app", apiCfg.MiddlewareMetricsInc(fs)))

	// API Endpoints
	mux.HandleFunc("GET /api/healthz", api.HandleCheckHealth)
	mux.HandleFunc("POST /api/users", apiCfg.HandleCreateUser)
	mux.HandleFunc("GET /api/chirps", apiCfg.HandleGetChirps)
	mux.HandleFunc("POST /api/chirps", apiCfg.HandleCreateChirp)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.HandleGetChirp)
	mux.HandleFunc("POST /api/login", apiCfg.HandleGetUser)

	// Administration
	mux.HandleFunc("GET /admin/metrics", apiCfg.HandleMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.HandleResetUsers)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	log.Printf("Server starting on %s...", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
