package router

// Logger defines the behaviour required of a logger needed by the router
type Logger interface {
	Errorw(string, ...interface{})
	Debugw(string, ...interface{})
}
