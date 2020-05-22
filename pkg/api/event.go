package api

import (
	"net/http"

	"github.com/skyerus/history-api/pkg/event/eventrepo"
	"github.com/skyerus/history-api/pkg/event/eventservice"
)

func (router router) randomHistoricalEvent(w http.ResponseWriter, r *http.Request) {
	eventRepo := eventrepo.NewEventRepo(router.db)
	eventService := eventservice.NewEventService(eventRepo)
	he, customErr := eventService.RandomHistoricalEvent()
	if customErr != nil {
		handleError(w, customErr)
		return
	}

	respondJSON(w, http.StatusOK, he)
}
