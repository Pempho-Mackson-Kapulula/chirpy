package api

import (
	"net/http"

	"github.com/Pempho-Mackson-Kapulula/chirpy/internal/auth"
)

func (cfg *Config) HandleRevokeRefreshToken(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't get token")
		return
	}

	err = cfg.DB.RevokeRefreshToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't revoke token")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
