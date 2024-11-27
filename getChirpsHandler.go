package main

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/google/uuid"
	"github.com/mmammel12/chirpy/internal/database"
)

func (cfg *apiConfig) getChirpsHandler(w http.ResponseWriter, r *http.Request) {
	sort := r.URL.Query().Get("sort")
	if sort == "" {
		sort = "asc"
	}

	if sort != "asc" && sort != "desc" {
		respondWithError(w, http.StatusBadRequest, "Invalid sort", fmt.Errorf("Invalid sort query param: %v", sort))
		return
	}

	authorIDQueryParam := r.URL.Query().Get("author_id")
	var chirps []database.Chirp
	var err error
	if authorIDQueryParam != "" {
		authorID, err := uuid.Parse(authorIDQueryParam)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author_id", err)
			return
		}
		chirps, err = cfg.db.ListChirpsByAuthor(r.Context(), authorID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps by author", err)
			return
		}
	} else {
		chirps, err = cfg.db.ListChirps(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
			return
		}
	}

	jsonChirps := []Chirp{}
	for _, dbChirp := range chirps {
		jsonChirps = append(jsonChirps, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			UserID:    dbChirp.UserID,
			Body:      dbChirp.Body,
		})
	}

	if sort == "desc" {
		slices.Reverse(jsonChirps)
	}

	respondWithJSON(w, http.StatusOK, jsonChirps)
}
