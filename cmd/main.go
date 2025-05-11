package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bonzonkim/play-with-gotth/internal/templates"
)

type Response struct {
	Message string
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/*", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Home().Render(r.Context(), w)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		res := Response{
			Message: "Hello, World",
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Printf("Failed to encode response %v", err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
