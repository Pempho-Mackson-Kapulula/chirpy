package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create the request router
	mux := http.NewServeMux()

	// Define the server configuration
	port := "8080"
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Create a handler that serves static files from the current dir
	fileServerHandler := http.FileServer(http.Dir("."))

	// Register the file servever handler at the root path
	mux.Handle("/", fileServerHandler)

	// Start the server and listen for incoming requests
	err := srv.ListenAndServe()

	// Check if the server failed to start (e.g., port is already in use)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
