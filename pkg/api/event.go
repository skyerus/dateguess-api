package api

import (
	"net/http"
	"sort"
	"strconv"

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

func (router router) randomHistoricalEvents(w http.ResponseWriter, r *http.Request) {
	qty, err := strconv.Atoi(r.URL.Query().Get("qty"))
	if err != nil {
		respondBadRequest(w)
		return
	}
	eventRepo := eventrepo.NewEventRepo(router.db)
	eventService := eventservice.NewEventService(eventRepo)
	hes, customErr := eventService.RandomHistoricalEvents(qty)
	if customErr != nil {
		handleError(w, customErr)
		return
	}
	if r.URL.Query().Get("order") == "true" {
		sort.Slice(*hes, func(i, j int) bool {
			return (*hes)[i].Date.Before((*hes)[j].Date)
		})
	}

	respondJSON(w, http.StatusOK, hes)
}
