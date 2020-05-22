package event

import (
	"time"

	"github.com/skyerus/history-api/pkg/customerror"
	"github.com/skyerus/history-api/pkg/models"
)

// Repository ...
type Repository interface {
	SaveHistoricalEvent(he *models.HistoricalEvent) customerror.Error
	SaveBirthEvent(be *models.BirthEvent) customerror.Error
	SaveDeathEvent(de *models.DeathEvent) customerror.Error
	SaveHolidayEvent(he *models.HolidayEvent) customerror.Error
	HistoricalEventIdsBetween(from time.Time, to time.Time) (*[]int, customerror.Error)
	GetHistoricalEvent(ID int) (*models.HistoricalEvent, customerror.Error)
}
