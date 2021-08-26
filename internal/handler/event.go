package handler

import (
	"dateguess-api/internal/model"
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

type Event struct {
	router router
	logger logger
	es     eventService
}

type eventService interface {
	RandomHistoricalEvent() (model.HistoricalEvent, error)
	RandomHistoricalEvents(qty int) ([]model.HistoricalEvent, error)
}

func NewEvent(r router, l logger, es eventService) *Event {
	return &Event{
		router: r,
		logger: l,
		es:     es,
	}
}

func (h *Event) RandomHistoricalEvent(w http.ResponseWriter, r *http.Request) {
	e, err := h.es.RandomHistoricalEvent()
	if err != nil {
		h.router.Error(r.Context(), w, http.StatusInternalServerError, "INTERNAL SERVER ERROR",
			fmt.Errorf("failed to get random historical event: %w", err),
		)
		return
	}

	h.router.Respond(r.Context(), w, http.StatusOK, e)
}

func (h *Event) RandomHistoricalEvents(w http.ResponseWriter, r *http.Request) {
	qty, err := strconv.Atoi(r.URL.Query().Get("qty"))
	if err != nil {
		h.router.Respond(r.Context(), w, http.StatusBadRequest, "qty must be integer")
		return
	}

	es, err := h.es.RandomHistoricalEvents(qty)
	if err != nil {
		h.router.Error(r.Context(), w, http.StatusInternalServerError, "INTERNAL SERVER ERROR",
			fmt.Errorf("failed to get random historical events: %w", err),
		)
		return
	}
	if r.URL.Query().Get("order") == "true" {
		sort.Slice(es, func(i, j int) bool {
			return (es)[i].Date.Before((es)[j].Date)

		})
	}

	h.router.Respond(r.Context(), w, http.StatusOK, es)
}
