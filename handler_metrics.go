package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Increment the atomic counter on every incoming request
		cfg.fileserverHits.Add(1)

		// Pass the request to the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// 1. Generate the formatted HTML string variable
	htmlResponse := fmt.Sprintf(`<html>
  	<body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  	</body>
	</html>`, cfg.fileserverHits.Load())

	// 2. Convert the string to bytes and write it to the response
	w.Write([]byte(htmlResponse))
}
