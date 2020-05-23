package eventservice

import (
	"math/rand"
	"time"

	"github.com/skyerus/history-api/pkg/customerror"
	"github.com/skyerus/history-api/pkg/event"
	"github.com/skyerus/history-api/pkg/models"
)

type eventService struct {
	eventRepo event.Repository
}

// NewEventService ...
func NewEventService(er event.Repository) event.Service {
	return &eventService{er}
}

func randomTime() time.Time {
	var start int64
	start = -62167176000 // Jan 1st 0000
	randomUnix := rand.Int63n(time.Now().Unix()-start) + start
	return time.Unix(randomUnix, 0)
}

func (es eventService) RandomHistoricalEvent() (*models.HistoricalEvent, customerror.Error) {
	randTime := randomTime()
	historicalEventIds, customErr := es.eventRepo.HistoricalEventIdsBetween(randTime, randTime.AddDate(0, 6, 0))
	if customErr != nil {
		return nil, customErr
	}
	numOfIds := len(*historicalEventIds)
	if numOfIds == 0 {
		return es.RandomHistoricalEvent()
	}
	randIndex := rand.Intn(numOfIds)
	randID := (*historicalEventIds)[randIndex]

	return es.eventRepo.GetHistoricalEvent(randID)
}

func (es eventService) RandomHistoricalEvents(qty int) (*[]models.HistoricalEvent, customerror.Error) {
	var historicalEvents []models.HistoricalEvent
	i := 0
	for i < qty {
		historicalEvent, customErr := es.RandomHistoricalEvent()
		if customErr != nil {
			return nil, customErr
		}
		for !eventUnique(historicalEvent, &historicalEvents) {
			historicalEvent, customErr = es.RandomHistoricalEvent()
			if customErr != nil {
				return nil, customErr
			}
		}
		historicalEvents = append(historicalEvents, *historicalEvent)
		i++
	}

	return &historicalEvents, nil
}

func eventUnique(he *models.HistoricalEvent, hes *[]models.HistoricalEvent) bool {
	for _, historicalEvent := range *hes {
		if he.ID == historicalEvent.ID {
			return false
		}
	}

	return true
}
