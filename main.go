package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

//go:embed templates/*
var templateFiles embed.FS

func main() {
	// Serve static files
	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))

	// Serve index page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := templateFiles.ReadFile("templates/index.html")
		if err != nil {
			http.Error(w, "Page not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(data)
	})

	port := ":8080"
	log.Printf("Server starting on http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
