package repository

// Logger defines the behaviour required of a logger
type Logger interface {
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
}
