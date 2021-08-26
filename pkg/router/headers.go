package router

import (
	"context"
	"net/http"
)

type headersSetContextKey int

const (
	// HeadersSetKey is used to store whether headers have been set per request
	HeadersSetKey headersSetContextKey = 0
)

// SetHeaders sets headers for the response. Using the returned context to
// subsequent calls to the Respond funcs will overwrite the default headers
// configured. Using without retrieving the context sets the passed headers
// in addition to the default headers
func (r *Router) SetHeaders(
	ctx context.Context, w http.ResponseWriter, h map[string][]string,
) context.Context {
	header := w.Header()
	for k, v := range h {
		header[k] = v
	}

	return context.WithValue(ctx, HeadersSetKey, true)
}

// setDefaultHeaders sets the default (most commonly used) set of headers.
// Is called for each request when no specific headers have been set using
// SetHeaders
// nolint:lll
func (r *Router) setDefaultHeaders(ctx context.Context, w http.ResponseWriter) context.Context {
	if r.HeadersSet(ctx) {
		return ctx
	}

	r.SetHeaders(ctx, w, r.defaultHeaders)

	return ctx
}

// SetHeader sets a single header. Default headers will still be set after use.
// If you wish for default headers not to be set use SetHeaders instead
func (r *Router) SetHeader(w http.ResponseWriter, k, v string) {
	w.Header().Set(k, v)
}

// HeadersSet determines whether headers have been set
// for this request
func (r *Router) HeadersSet(ctx context.Context) bool {
	if set, ok := ctx.Value(HeadersSetKey).(bool); ok {
		return set
	}

	return false
}
