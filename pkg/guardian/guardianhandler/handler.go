package guardianhandler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/skyerus/history-api/pkg/customerror"
	"github.com/skyerus/history-api/pkg/handler"
	"github.com/skyerus/history-api/pkg/logger"
)

type guardianHandler struct {
}

// NewGuardianHandler ...
func NewGuardianHandler() handler.RequestHandler {
	return &guardianHandler{}
}

// SendRequest ...
func (w guardianHandler) SendRequest(request *http.Request) (*http.Response, customerror.Error) {
	url := request.URL
	q := url.Query()
	q.Set("api-key", os.Getenv("GUARDIAN_KEY"))
	q.Set("show-fields", "bodyText")
	request.Header.Add("Content-Type", "application/json")

	request.URL.RawQuery = q.Encode()

	var response *http.Response
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return response, customerror.NewGenericHTTPError(err)
	}
	if response.StatusCode == http.StatusUnauthorized {
		return response, customerror.NewUnauthorizedError(nil)
	}

	if response.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return response, customerror.NewGenericHTTPError(err)
		}
		bodyString := string(bodyBytes)
		go logger.Log(errors.New(bodyString))
		return response, customerror.NewGenericHTTPError(nil)
	}

	return response, nil
}
