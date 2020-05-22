package eventrepo

import (
	"database/sql"

	"github.com/skyerus/history-api/pkg/customerror"
	"github.com/skyerus/history-api/pkg/event"
	"github.com/skyerus/history-api/pkg/models"
)

type eventRepo struct {
	db *sql.DB
}

// NewEventRepo ...
func NewEventRepo(db *sql.DB) event.Repository {
	return &eventRepo{db}
}

func (er eventRepo) SaveHistoricalEvent(he *models.HistoricalEvent) customerror.Error {
	stmtIns, err := er.db.Prepare("INSERT INTO historical_event (date_time, fact) VALUES(?, ?)")
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(he.Date, he.Fact)
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}

	return nil
}

func (er eventRepo) SaveBirthEvent(be *models.BirthEvent) customerror.Error {
	stmtIns, err := er.db.Prepare("INSERT INTO birth_event (date_time, fact) VALUES(?, ?)")
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(be.Date, be.Fact)
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}

	return nil
}

func (er eventRepo) SaveDeathEvent(de *models.DeathEvent) customerror.Error {
	stmtIns, err := er.db.Prepare("INSERT INTO death_event (date_time, fact) VALUES(?, ?)")
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(de.Date, de.Fact)
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}

	return nil
}

func (er eventRepo) SaveHolidayEvent(he *models.HolidayEvent) customerror.Error {
	stmtIns, err := er.db.Prepare("INSERT INTO holiday_event (date_time, fact) VALUES(?, ?)")
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(he.Date, he.Fact)
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}

	return nil
}
