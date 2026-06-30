package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/Pempho-Mackson-Kapulula/chirpy/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// api config type
type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
}

func main() {
	//load .env file and get dbUrl
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")

	// prepare db connection pool
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	//create a database wrapper
	dbQueries := database.New(db)

	//create an instance of apiConfig
	apiCfg := apiConfig{
		db: dbQueries,
	}

	//router
	mux := http.NewServeMux()

	//server config
	addr := ":8080"
	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// static files server
	fs := http.FileServer(http.Dir("."))
	mux.Handle("/app/", http.StripPrefix("/app", apiCfg.middlewareMetricsInc(fs)))

	// health check
	mux.HandleFunc("GET /api/healthz", handlerHealthCheck)

	// metrics
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	//chirp validation
	mux.HandleFunc("POST /api/validate_chirp", handlerValidate)

	log.Fatal(srv.ListenAndServe())
}
