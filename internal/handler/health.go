package handler

import (
	"net/http"
)

// Health is the handler at /v1/health for checking service health
type Health struct {
	router router
}

// NewHealth returns a new Health handler
func NewHealth(r router) *Health {
	return &Health{
		router: r,
	}
}

// Get is the handlerfunc for GET /v1/health
func (h *Health) Get(w http.ResponseWriter, r *http.Request) {
	h.router.WriteHeader(r.Context(), w, http.StatusOK)
}
