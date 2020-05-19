package handler

import (
	"net/http"

	"github.com/skyerus/history-api/pkg/customerror"
)

// RequestHandler ...
type RequestHandler interface {
	SendRequest(request *http.Request) (*http.Response, customerror.Error)
}
