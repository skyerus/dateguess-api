// Package router holds logic for responding to HTTP requests
package router

// Router deals with responding to http requests
type Router struct {
	defaultHeaders map[string][]string
	logger         Logger
	reqIDCtxKey    interface{}
}

// New returns a new Router
func New(cfg Config, l Logger) *Router {
	return &Router{
		defaultHeaders: cfg.DefaultHeaders,
		logger:         l,
		reqIDCtxKey:    cfg.RequestIDContextKey,
	}
}
