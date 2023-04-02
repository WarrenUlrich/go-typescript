package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:generate npm --prefix ./client run build
//go:embed client/dist/*
var static embed.FS

func main() {
	mux := http.NewServeMux()

	// Serve static files
	staticHandler, err := setupStaticFileServer()
	if err != nil {
		panic(err)
	}
	mux.Handle("/", staticHandler)

	// API routes
	mux.HandleFunc("/api/hello", helloHandler)

	// Start the server
	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func setupStaticFileServer() (http.Handler, error) {
	// Get the embedded files as an fs.FS
	content, err := fs.Sub(static, "client/dist")
	if err != nil {
		return nil, err
	}

	// Create a file server for serving the static files
	return http.FileServer(http.FS(content)), nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	// Write a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"message": "Hello from the API!"}`)
}
