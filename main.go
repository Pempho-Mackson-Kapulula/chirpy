package main

import (
	"fmt"
	"net/http"
)

func main() {
	//creates the request router
	mux := http.NewServeMux()

	//defines the server configuration
	port := "8080"
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	//creates a handler that serves static files from the path
	fileServerHandler := http.FileServer(http.Dir("."))

	//registers the file server handler at the /app/ path
	mux.Handle("/app/", http.StripPrefix("/app", fileServerHandler))

	//registers the handler function for the /healthz route.
	mux.HandleFunc("/healthz", readinessHandler)

	//starts the server and listen for incoming requests
	err := srv.ListenAndServe()

	// checks if the server failed to start (e.g., port is already in use)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
