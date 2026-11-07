package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc("GET /stream/{track_id}", func(w http.ResponseWriter, r *http.Request) {
		trackID := r.PathValue("track_id")
		filePath := fmt.Sprintf("./storage/%s.mp3", trackID)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.Error(w, "Track not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "audio/mpeg")
		http.ServeFile(w, r, filePath)
	})

	addr := ":8080"
	log.Printf("Sanctum audio core listening on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
