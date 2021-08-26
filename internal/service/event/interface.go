package event

import (
	"dateguess-api/internal/model"
	"time"
)

type repository interface {
	SaveHistoricalEvent(e *model.HistoricalEvent) error
	SaveBirthEvent(e *model.BirthEvent) error
	SaveDeathEvent(e *model.DeathEvent) error
	SaveHolidayEvent(e *model.HolidayEvent) error
	HistoricalEventIdsBetween(from time.Time, to time.Time) ([]int, error)
	HistoricalEvent(id int) (model.HistoricalEvent, error)
}
