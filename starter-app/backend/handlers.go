package main

import (
	"encoding/json"
	"net/http"
)

// writeJSON is a small helper: set the header, the status code, then encode.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// GET /api/health  → {"status":"ok"}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// GET /api/todos  → [ {...}, {...} ]
func (s *Store) listTodos(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	writeJSON(w, http.StatusOK, s.todos)
}

// ---------------------------------------------------------------------------
// YOUR TURN (Step 3). See docs/API.md for the exact request/response contract.
// ---------------------------------------------------------------------------

// GET /api/todos/{id}
// Hints: r.PathValue("id") gives the id as a string → strconv.Atoi.
//
//	Not found → 404 with {"error":"todo not found"}.
func (s *Store) getTodo(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, map[string]string{"error": "not implemented"})
}

// POST /api/todos   body: {"title":"..."}
// Hints: json.NewDecoder(r.Body).Decode(&input)
//
//	Validate: empty title → 400. Success → 201 + the new todo.
func (s *Store) createTodo(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, map[string]string{"error": "not implemented"})
}

// PATCH /api/todos/{id}   body: {"title":"...", "done":true}  (both optional)
// Hint: use pointer fields (*string, *bool) in your input struct so you can
//
//	tell "not sent" (nil) apart from "sent as false/empty".
func (s *Store) updateTodo(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, map[string]string{"error": "not implemented"})
}

// DELETE /api/todos/{id}  → 204 No Content (empty body)
func (s *Store) deleteTodo(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotImplemented, map[string]string{"error": "not implemented"})
}
