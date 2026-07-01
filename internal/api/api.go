package api

import (
	"sync/atomic"

	"github.com/Pempho-Mackson-Kapulula/chirpy/internal/database"
)

type Config struct {
	fileserverHits atomic.Int32
	DB             *database.Queries
	Platform       string
}
