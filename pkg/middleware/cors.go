package middleware

import (
	"net/http"
)

// Cors is middleware for cors
func (mw *Middleware) Cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if mw.dev {
			w.Header().Set("Access-Control-Allow-Origin", mw.allowOrigin)
		}
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Set-Cookie")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
