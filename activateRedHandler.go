package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/mmammel12/chirpy/internal/auth"
	"github.com/mmammel12/chirpy/internal/database"
)

func (cfg *apiConfig) activateRedHandler(w http.ResponseWriter, r *http.Request) {
	type body struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	key, err := auth.GetAPIKey(r.Header)
	if err != nil || key != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "Invalid API Key", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := body{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode body", err)
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err = cfg.db.UpdateChirypRedStatus(r.Context(), database.UpdateChirypRedStatusParams{
		ID:          params.Data.UserID,
		IsChirpyRed: true,
	})
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Could not find user", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
