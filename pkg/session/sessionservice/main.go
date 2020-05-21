package sessionservice

import (
	"time"

	"github.com/skyerus/history-api/pkg/customerror"
	"github.com/skyerus/history-api/pkg/session"
	"github.com/skyerus/ims-api/pkg/logger"
)

type sessionService struct {
	sessionRepo session.Repository
}

// NewSessionService ...
func NewSessionService(sr session.Repository) session.Service {
	return &sessionService{sr}
}

func (ss sessionService) LogSession(ipAddress string) customerror.Error {
	now := time.Now()
	customErr := ss.sessionRepo.LogSession(ipAddress, now)
	if customErr != nil {
		logger.Log(customErr.OriginalError())
		return customErr
	}

	return nil
}
