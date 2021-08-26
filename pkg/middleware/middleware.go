package middleware

// Middleware hosts the middleware methods
type Middleware struct {
	allowOrigin string
	dev         bool
}

// NewMiddleware constructs a new Middleware
func NewMiddleware(allowOrigin string, dev bool) *Middleware {
	return &Middleware{
		allowOrigin: allowOrigin,
		dev:         dev,
	}
}
