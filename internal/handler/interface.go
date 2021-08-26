package handler

import (
	"context"
	"net/http"
)

// router defines the behaviour of the router needed by the http handlers
type router interface {
	Respond(context.Context, http.ResponseWriter, int, interface{})
	RespondRaw(context.Context, http.ResponseWriter, int, []byte)
	Error(context.Context, http.ResponseWriter, int, interface{}, interface{})
	Success(context.Context, http.ResponseWriter, int)
	ServeFile(context.Context, http.ResponseWriter, *http.Request, string)
	SetHeaders(context.Context, http.ResponseWriter, map[string][]string) context.Context
	WriteHeader(context.Context, http.ResponseWriter, int)
}

// logger defines the behaviour required of a logger needed by the http
// handlers
type logger interface {
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
}

type statusCoder interface {
	StatusCode() int
}

type conflicter interface {
	Conflict() bool
}
