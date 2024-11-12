package main

import (
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

func (cfg *apiConfig) fileserverHitsHandler(res http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.New("metrics").Parse(`
        <html>
            <body>
                <h1>Welcome, Chirpy Admin</h1>
                <p>Chirpy has been visited {{.}} times!</p>
            </body>
        </html>
        `)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(res, cfg.fileserverHits.Load())
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (cfg *apiConfig) resetMetricsHandler(res http.ResponseWriter, _ *http.Request) {
	cfg.fileserverHits.Store(0)

	res.WriteHeader(http.StatusOK)
}
