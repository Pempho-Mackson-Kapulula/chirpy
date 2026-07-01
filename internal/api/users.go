package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *Config) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Email string `json:"email"`
	}

	//decode request email
	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		return
	}

	type User struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email     string    `json:"email"`
	}

	respondWithJSON(w, http.StatusCreated, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})

}

func (cfg *Config) HandleResetUsers(w http.ResponseWriter, r *http.Request) {
	if cfg.Platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}

	err := cfg.DB.DeleteAllUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong on our end")
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

}
