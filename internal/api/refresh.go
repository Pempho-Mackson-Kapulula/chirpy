package api

import (
	"net/http"
	"time"

	"github.com/Pempho-Mackson-Kapulula/chirpy/internal/auth"
)

func (cfg *Config) HandleRefreshToken(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't get token")
		return
	}

	userID, err := cfg.DB.GetUserFromRefreshToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid token")
		return
	}

	accessToken, err := auth.MakeJWT(userID, cfg.Secret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create jwt")
		return
	}

	type Response struct {
		Token string `json:"token"`
	}

	respondWithJSON(w, http.StatusOK, Response{
		Token: accessToken,
	})
}
