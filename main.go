package main

import (
	"net/http"
	"github.com/admarcel/go_http_server/version"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/admarcel/go_http_server/handlers"
	"fmt"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/version", version.Handler)
	http.HandleFunc("/", handlers.RootHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Listening on :8000")
}
