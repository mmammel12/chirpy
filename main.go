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

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/healthz", healthCheckHandler)

	fmt.Printf("server listening for requests on port %v\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
