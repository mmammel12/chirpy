package main

import (
	"sync/atomic"

	"github.com/mmammel12/chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
	platform       string
	jwtSecret      string
}
