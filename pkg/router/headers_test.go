package router_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"dateguess-api/pkg/router"

	"go.uber.org/zap"
)

func TestSetHeadersOverWrite(t *testing.T) {
	r := router.New(
		router.Config{
			DefaultHeaders: map[string][]string{
				"X-Request-ID": {"1234-abcd"},
			},
		},
		zap.NewExample().Sugar(),
	)

	w := httptest.NewRecorder()

	headers := map[string][]string{"Expected": {"Header"}}

	payload := testStruct{
		Name: "hi",
		Age:  0,
	}

	ctx := r.SetHeaders(context.Background(), w, headers)
	r.Respond(ctx, w, http.StatusOK, payload)

	res := w.Result()
	defer res.Body.Close()

	got := res.Header

	expected := map[string][]string{
		"Content-Type": {"application/json"},
		"Expected":     {"Header"},
	}

	equal(t, (map[string][]string)(got), expected)
}

func TestSetHeadersAppend(t *testing.T) {
	r := router.New(
		router.Config{
			DefaultHeaders: map[string][]string{
				"X-Request-ID": {"1234-abcd"},
			},
		},
		zap.NewExample().Sugar(),
	)

	w := httptest.NewRecorder()

	headers := map[string][]string{"Expected": {"Header"}}

	payload := testStruct{
		Name: "hi",
		Age:  0,
	}

	ctx := context.Background()
	r.SetHeaders(ctx, w, headers)
	r.Respond(ctx, w, http.StatusOK, payload)

	res := w.Result()
	defer res.Body.Close()

	got := res.Header

	expected := map[string][]string{
		"Content-Type": {"application/json"},
		"Expected":     {"Header"},
		"X-Request-ID": {"1234-abcd"},
	}

	equal(t, (map[string][]string)(got), expected)
}
