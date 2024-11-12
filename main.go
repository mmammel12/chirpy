// Package main
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	apiCfg := apiConfig{}

	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	mux.HandleFunc("GET /admin/metrics", apiCfg.fileserverHitsHandler)
	mux.HandleFunc("GET /api/healthz", healthCheckHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.resetMetricsHandler)

	fmt.Printf("server listening for requests on port %v\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
