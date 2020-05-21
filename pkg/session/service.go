package session

import "github.com/skyerus/history-api/pkg/customerror"

// Service ...
type Service interface {
	LogSession(ipAddress string) customerror.Error
}
