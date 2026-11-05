// Entry point of the Go backend.
// Run with:  go run .
// Then open: http://localhost:8080/api/health
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	store := NewStore() // in-memory data for now (Step 6 replaces it with a database)

	mux := http.NewServeMux()

	// --- Routes -------------------------------------------------------------
	// Go 1.22+ lets you put the HTTP method in the pattern: "GET /path".
	mux.HandleFunc("GET /api/health", healthHandler)
	mux.HandleFunc("GET /api/todos", store.listTodos)

	// TODO (Step 3): uncomment these when you implement them in handlers.go
	// mux.HandleFunc("GET /api/todos/{id}", store.getTodo)
	// mux.HandleFunc("POST /api/todos", store.createTodo)
	// mux.HandleFunc("PATCH /api/todos/{id}", store.updateTodo)
	// mux.HandleFunc("DELETE /api/todos/{id}", store.deleteTodo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("backend listening on http://localhost:%s", port)
	// Every request passes through the CORS middleware before reaching mux.
	log.Fatal(http.ListenAndServe(":"+port, withCORS(mux)))
}
