package main

import "net/http"

func healthCheckHandler(res http.ResponseWriter, _ *http.Request) {
	res.Header().Add("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("OK"))
}
