package main

import (
	"encoding/json"
	"fmt"
	"github.com/admarcel/go_http_server/version"
	"net/http"
)

func main() {

	version := version.GetVersion()
	fmt.Printf("version: %s", version.String())
	fmt.Printf("version: %s", version.String())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to version %s", version.String())
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		payload, err := json.Marshal(version)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
