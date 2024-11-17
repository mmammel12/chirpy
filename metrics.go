package main

import (
	"errors"
	"html/template"
	"net/http"
)

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	},
	)
}

func (cfg *apiConfig) fileserverHitsHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.New("metrics").Parse(`
        <html>
            <body>
                <h1>Welcome, Chirpy Admin</h1>
                <p>Chirpy has been visited {{.}} times!</p>
            </body>
        </html>
        `)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, cfg.fileserverHits.Load())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (cfg *apiConfig) resetMetricsHandler(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Forbidden", errors.New("Forbidden platform"))
		return
	}

	cfg.fileserverHits.Store(0)
	err := cfg.db.DeleteUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not delete users", err)
		return
	}
	err = cfg.db.DeleteChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not delete chirps", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
