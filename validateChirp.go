package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func validateChirp(w http.ResponseWriter, r *http.Request) {
	type Chirp struct {
		Body string `json:"body"`
	}

	type Err struct {
		Error string `json:"error"`
	}

	w.Header().Set("Content-Type", "appication/json")

	decoder := json.NewDecoder(r.Body)
	chirp := Chirp{}
	err := decoder.Decode(&chirp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Err{Error: "Something went wrong"})
		return
	}

	if len(chirp.Body) > 140 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Err{Error: "Chirp is too long"})
		return
	}

	badWords := map[string]bool{
		"kerfuffle": true,
		"sharbert":  true,
		"fornax":    true,
	}

	splitBody := strings.Split(chirp.Body, " ")
	for i, word := range splitBody {
		if _, exists := badWords[strings.ToLower(word)]; exists {
			splitBody[i] = "****"
		}
	}

	resBody := strings.Join(splitBody, " ")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		CleanedBody string `json:"cleaned_body"`
	}{CleanedBody: resBody})
}
