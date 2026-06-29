package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

// api config type
type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	//router
	mux := http.NewServeMux()

	//create an instance of apiConfig
	apiCfg := apiConfig{}

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
	mux.HandleFunc("GET /healthz", handlerHealthCheck)

	// metrics
	mux.HandleFunc("GET /metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /reset", apiCfg.handlerReset)

	log.Fatal(srv.ListenAndServe())
}
