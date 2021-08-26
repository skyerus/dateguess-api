package repository

// ConflictError implements the conflicter interface
type ConflictError struct {
	IsConflict bool
}

// Conflict reports a conflict
func (e ConflictError) Conflict() bool {
	return e.IsConflict
}

// Error implements the error interface
func (e ConflictError) Error() string {
	return "there is a conflict"
}

// StatusError provides a status code alongside an error
type StatusError struct {
	Err      error
	HTTPCode int
}

func (e StatusError) Error() string {
	return e.Err.Error()
}

// StatusCode returns the status code
func (e StatusError) StatusCode() int {
	return e.HTTPCode
}
