package main

import "net/http"

// withCORS lets the Next.js app (a different origin: localhost:3000) call this
// API from the browser. Without it the browser blocks the response.
// Read more: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Browsers send an OPTIONS "preflight" before POST/PATCH/DELETE with JSON.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// TODO (Step 7): write a withLogging middleware that prints
// method, path, and how long the request took. Chain it like:
//   withLogging(withCORS(mux))
