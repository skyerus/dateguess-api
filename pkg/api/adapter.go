package api

import (
	"net/http"
	"os"

	"github.com/skyerus/history-api/pkg/session/sessionrepo"
	"github.com/skyerus/history-api/pkg/session/sessionservice"
)

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("ENV") == "dev" {
			w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_ALLOW_ORIGIN"))
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

func (router router) handleSession(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			cookie = &http.Cookie{Name: "session", Value: r.RemoteAddr, Domain: os.Getenv("API_DOMAIN"), MaxAge: 7200, Path: "/"}
			sessionRepo := sessionrepo.NewSessionRepo(router.db)
			sessionService := sessionservice.NewSessionService(sessionRepo)
			go sessionService.LogSession(r.RemoteAddr)
			http.SetCookie(w, cookie)
		}
		h.ServeHTTP(w, r)
	})
}
