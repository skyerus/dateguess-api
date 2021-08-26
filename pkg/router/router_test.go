package router_test

import (
	"context"
	"dateguess-api/pkg/router"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
)

type testStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func equal(t *testing.T, got, expected interface{}) {
	t.Helper()

	if diff := cmp.Diff(got, expected); diff != "" {
		t.Errorf(
			"Wrong results (-got +expected)\n%s\n",
			diff,
		)
	}
}

func TestRespond(t *testing.T) {
	r := router.New(
		router.Config{
			DefaultHeaders: map[string][]string{
				"Content-Type": {"application/json"},
			},
		},
		zap.NewExample().Sugar(),
	)

	w := httptest.NewRecorder()

	payload := testStruct{
		Name: "John",
		Age:  99,
	}

	r.Respond(context.Background(), w, http.StatusOK, payload)

	res := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(res.Body)

	var got testStruct

	if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
		t.Error(err)
	}

	equal(t, res.StatusCode, http.StatusOK)
	equal(t, got, payload)
}

func TestRespondRaw(t *testing.T) {
	r := router.New(
		router.Config{
			DefaultHeaders: map[string][]string{
				"Content-Type": {"application/json"},
			},
		},
		zap.NewExample().Sugar(),
	)

	w := httptest.NewRecorder()

	payload := []byte("I am a payload")

	r.RespondRaw(context.Background(), w, http.StatusOK, payload)

	res := w.Result()
	defer res.Body.Close()

	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	equal(t, res.StatusCode, http.StatusOK)
	equal(t, got, payload)
}

func TestSuccess(t *testing.T) {
	r := router.New(
		router.Config{
			DefaultHeaders: map[string][]string{
				"Content-Type": {"application/json"},
			},
		},
		zap.NewExample().Sugar(),
	)

	w := httptest.NewRecorder()

	r.Success(context.Background(), w, http.StatusOK)

	res := w.Result()
	defer res.Body.Close()

	equal(t, res.StatusCode, http.StatusOK)
}

func TestError(t *testing.T) {
	r := router.New(
		router.Config{
			DefaultHeaders: map[string][]string{
				"Content-Type": {"application/json"},
			},
		},
		zap.NewExample().Sugar(),
	)

	w := httptest.NewRecorder()

	err := errors.New("I am an error message")

	expected := router.ErrorResponse{
		Error: err.Error(),
	}

	r.Error(context.Background(), w, http.StatusBadRequest, err, err)

	res := w.Result()
	defer res.Body.Close()

	var got router.ErrorResponse

	if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
		t.Error(err)
	}

	equal(t, res.StatusCode, http.StatusBadRequest)
	equal(t, got, expected)
}
