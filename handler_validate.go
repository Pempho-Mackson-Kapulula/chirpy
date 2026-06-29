package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type returnvals struct {
		Error string `json:"error"`
	}
	respBody := returnvals{
		Error: msg,
	}
	respondWithJSON(w, code, respBody)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

var badWords = map[string]struct{}{
	"kerfuffle": {},
	"sharbert":  {},
	"fornax":    {},
}

func cleanProfanity(text string) string {
	words := strings.Fields(text)
	for i, word := range words {
		lowerWord := strings.ToLower(word)
		if _, exists := badWords[lowerWord]; exists {
			words[i] = "****"
		}
	}

	return strings.Join(words, " ")
}

func handlerValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	//decode request body
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	cleanedBody := cleanProfanity(params.Body)

	type validResponse struct {
		CleanedBody string `json:"cleaned_body"`
	}

	respondWithJSON(w, 200, validResponse{
		CleanedBody: cleanedBody,
	})
}
