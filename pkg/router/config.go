package router

// Config holds the config for the router
type Config struct {
	DefaultHeaders      map[string][]string `json:"default-headers"`
	RequestIDContextKey interface{}         `json:"-"`
}
