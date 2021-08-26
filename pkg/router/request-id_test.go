package router_test

import (
	"context"
	"dateguess-api/pkg/router"
	"testing"

	"go.uber.org/zap"
)

type testReqIDKey int

const (
	reqIDKey testReqIDKey = 0
)

func TestGetRequestID(t *testing.T) {
	r := router.New(
		router.Config{
			DefaultHeaders: map[string][]string{
				"Content-Type": {"text/xml"},
			},
			RequestIDContextKey: reqIDKey,
		},
		zap.NewExample().Sugar(),
	)

	ctx := context.WithValue(
		context.Background(),
		reqIDKey,
		"12345-abcde",
	)

	id := r.GetRequestID(ctx)

	equal(t, id, "12345-abcde")
}
