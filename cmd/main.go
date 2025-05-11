package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bonzonkim/play-with-gotth/internal/templates"
)

type Response struct {
	Message string
}

func main() {
	// static file serving
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		// Extract file path
		path := filepath.Clean(r.URL.Path)
		// set header text/css if file is css file
		if filepath.Ext(path) == ".css" {
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
		}
		fs := http.FileServer(http.Dir("static"))
		http.StripPrefix("/static/", fs).ServeHTTP(w, r)
	})

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

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
