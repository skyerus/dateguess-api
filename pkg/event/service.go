package event

import (
	"github.com/skyerus/history-api/pkg/customerror"
	"github.com/skyerus/history-api/pkg/models"
)

// Service ...
type Service interface {
	RandomHistoricalEvent() (*models.HistoricalEvent, customerror.Error)
	RandomHistoricalEvents(qty int) (*[]models.HistoricalEvent, customerror.Error)
}
