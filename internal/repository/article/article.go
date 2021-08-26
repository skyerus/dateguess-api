package article

import (
	"dateguess-api/internal/model"
	"dateguess-api/internal/repository"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type KeyValue struct {
	Key   string
	Value string
}

type Repository struct {
	apiKey  string
	baseURL string
	logger  repository.Logger
}

const searchEndpoint string = "/search"

func NewRepository(apiKey string, baseURL string, logger repository.Logger) *Repository {
	return &Repository{
		apiKey:  apiKey,
		baseURL: baseURL,
		logger:  logger,
	}
}

func (r *Repository) Search(params model.SearchParams) (model.GuardianContent, error) {
	var guardianContent model.GuardianContent
	request, err := http.NewRequest(
		"GET", fmt.Sprintf("%s%s", r.baseURL, searchEndpoint), nil,
	)
	if err != nil {
		return guardianContent, err
	}
	q := request.URL.Query()
	q.Add("page-size", strconv.Itoa(params.PageSize))
	q.Add("page", strconv.Itoa(params.Page))
	q.Add("section", params.Section)
	q.Add("from-date", params.From.Format("2006-01-02"))
	q.Add("to-date", params.To.Format("2006-01-02"))
	request.URL.RawQuery = q.Encode()

	response, err := r.sendRequest(request)
	if err != nil {
		return guardianContent, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			r.logger.Errorf("failed to close body: %v", err)
		}
	}(response.Body)

	if err := json.NewDecoder(response.Body).Decode(&guardianContent); err != nil {
		return guardianContent, fmt.Errorf("failed to decode json: %w", err)
	}

	return guardianContent, nil
}

func (r *Repository) sendRequest(request *http.Request) (*http.Response, error) {
	url := request.URL
	q := url.Query()
	q.Set("api-key", r.apiKey)
	q.Set("show-fields", "bodyText")
	request.Header.Add("Content-Type", "application/json")

	request.URL.RawQuery = q.Encode()

	var response *http.Response
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return response, fmt.Errorf("failed to execute request: %w", err)
	}
	if response.StatusCode == http.StatusUnauthorized {
		return response, errors.New("unauthorized")
	}

	if response.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return response, fmt.Errorf("failed to read response: %w", err)
		}
		bodyString := string(bodyBytes)

		return response, fmt.Errorf(
			"unexpected status code: %d, body: %s",
			response.StatusCode,
			bodyString,
		)
	}

	return response, nil
}
