package api

import (
	"net/http"
)

func (cfg *Config) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Increment the atomic counter on every incoming request
		cfg.fileserverHits.Add(1)

		// Pass the request to the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
