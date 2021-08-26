package router

import (
	"context"
)

// GetRequestID returns a requestID from the context
func (r *Router) GetRequestID(ctx context.Context) string {
	if r.reqIDCtxKey == nil {
		return "-"
	}

	if id, ok := ctx.Value(r.reqIDCtxKey).(string); ok {
		return id
	}

	return "-"
}
