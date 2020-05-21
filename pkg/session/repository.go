package session

import (
	"time"

	"github.com/skyerus/history-api/pkg/customerror"
)

// Repository ...
type Repository interface {
	LogSession(ipAddress string, t time.Time) customerror.Error
}
