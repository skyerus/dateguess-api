package sessionrepo

import (
	"database/sql"
	"time"

	"github.com/skyerus/history-api/pkg/customerror"
	"github.com/skyerus/history-api/pkg/session"
)

type sessionRepo struct {
	db *sql.DB
}

// NewSessionRepo ...
func NewSessionRepo(conn *sql.DB) session.Repository {
	return &sessionRepo{conn}
}

func (sr sessionRepo) LogSession(ipAddress string, t time.Time) customerror.Error {
	stmtIns, err := sr.db.Prepare("INSERT INTO session (ip, date_time) VALUES(?, ?)")
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(ipAddress, t)
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}

	return nil
}
