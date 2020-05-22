package eventrepo

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

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

func (er eventRepo) HistoricalEventIdsBetween(from time.Time, to time.Time) (*[]int, customerror.Error) {
	var ids []int
	var err error
	results, err := er.db.Query("SELECT id FROM historical_event WHERE `date_time` BETWEEN '" + from.Format("2006-01-01") + "' AND '" + to.Format("2006-01-01") + "';")
	if err != nil {
		return nil, customerror.NewGenericHTTPError(err)
	}
	defer results.Close()

	for results.Next() {
		var id int
		err = results.Scan(&id)
		if err != nil {
			return nil, customerror.NewGenericHTTPError(err)
		}
		ids = append(ids, id)
	}

	return &ids, nil
}

func (er eventRepo) GetHistoricalEvent(ID int) (*models.HistoricalEvent, customerror.Error) {
	results, err := er.db.Query("SELECT date_time, fact FROM historical_event WHERE id=" + strconv.Itoa(ID))
	if err != nil {
		return nil, customerror.NewGenericHTTPError(err)
	}
	defer results.Close()
	res := results.Next()
	if !res {
		return nil, customerror.NewGenericHTTPError(errors.New("No event found with that ID"))
	}
	var he models.HistoricalEvent
	he.ID = ID
	err = results.Scan(&he.Date, &he.Fact)
	if err != nil {
		return nil, customerror.NewGenericHTTPError(err)
	}

	return &he, nil
}
