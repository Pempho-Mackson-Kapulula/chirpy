package main

import (
	"log"
	"net/http"
)

func main() {
	//router
	mux := http.NewServeMux()

	//server
	addr := ":8080"
	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// static files server
	fs := http.FileServer(http.Dir("."))
	mux.Handle("/app/", http.StripPrefix("/app", fs))

	// health check
	mux.HandleFunc("/healthz", healthCheckHandler)

	log.Fatal(srv.ListenAndServe())
}
