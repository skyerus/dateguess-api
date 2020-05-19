package api

import (
	"encoding/json"
	"net/http"

	"github.com/skyerus/history-api/pkg/customerror"
	"github.com/skyerus/history-api/pkg/logger"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"message": message})
}

func respondGenericError(w http.ResponseWriter, err error) {
	go logger.Log(err)
	respondJSON(w, http.StatusInternalServerError, map[string]string{"message": "Oops, something went wrong. Please try again later."})
}

func respondBadRequest(w http.ResponseWriter) {
	respondJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid request"})
}

func respondUnauthorizedRequest(w http.ResponseWriter) {
	respondJSON(w, http.StatusUnauthorized, map[string]string{"message": "Unauthorized request"})
}

func handleError(w http.ResponseWriter, customerror customerror.Error) {
	if customerror.OriginalError() != nil {
		go logger.Log(customerror.OriginalError())
	}
	respondError(w, customerror.Code(), customerror.Message())
}
