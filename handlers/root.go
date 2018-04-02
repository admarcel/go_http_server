package handlers

import (
	"net/http"
	"encoding/json"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: "Hello world",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(payload)
}
