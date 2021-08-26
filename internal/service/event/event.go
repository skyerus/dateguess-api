package event

import (
	"dateguess-api/internal/model"
	rand2 "dateguess-api/pkg/rand"
	"math/rand"
)

// Service provides an API for performing actions on events
type Service struct {
	repo repository
}

// NewService constructs a Service
func NewService(repo repository) *Service {
	return &Service{repo}
}

func (s *Service) RandomHistoricalEvent() (model.HistoricalEvent, error) {
	randTime := rand2.Time()
	historicalEventIds, err := s.repo.HistoricalEventIdsBetween(
		randTime, randTime.AddDate(0, 6, 0),
	)
	if err != nil {
		return model.HistoricalEvent{}, err
	}
	if len(historicalEventIds) == 0 {
		return s.RandomHistoricalEvent()
	}
	randIndex := rand.Intn(len(historicalEventIds))
	randID := historicalEventIds[randIndex]

	return s.repo.HistoricalEvent(randID)
}

func (s *Service) RandomHistoricalEvents(qty int) ([]model.HistoricalEvent, error) {
	var historicalEvents []model.HistoricalEvent

	for i := 0; i < qty; i++ {
		historicalEvent, err := s.RandomHistoricalEvent()
		if err != nil {
			return nil, err
		}
		for !isEventUnique(historicalEvent, &historicalEvents) {
			historicalEvent, err = s.RandomHistoricalEvent()
			if err != nil {
				return nil, err
			}
		}
		historicalEvents = append(historicalEvents, historicalEvent)
	}

	return historicalEvents, nil
}

func isEventUnique(he model.HistoricalEvent, hes *[]model.HistoricalEvent) bool {
	for _, historicalEvent := range *hes {
		if he.ID == historicalEvent.ID {
			return false
		}
	}

	return true
}
