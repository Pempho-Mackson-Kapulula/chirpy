package api

import (
	"fmt"
	"net/http"
)

func (cfg *Config) HandleMetrics(w http.ResponseWriter, r *http.Request) {
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
	if _, err := w.Write([]byte(htmlResponse)); err != nil {
		return
	}
}
