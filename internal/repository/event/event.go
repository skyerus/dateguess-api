package event

import (
	"database/sql"
	"dateguess-api/internal/model"
	"dateguess-api/internal/repository"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Repository provides the data layer
type Repository struct {
	db     *sql.DB
	logger repository.Logger
}

// NewRepository constructs a Repository
func NewRepository(db *sql.DB, l repository.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: l,
	}
}

func (r *Repository) SaveHistoricalEvent(e *model.HistoricalEvent) error {
	return r.saveEvent("historical_event", e.Date, e.Fact)
}

func (r *Repository) SaveBirthEvent(e *model.BirthEvent) error {
	return r.saveEvent("birth_event", e.Date, e.Fact)
}

func (r *Repository) SaveDeathEvent(e *model.DeathEvent) error {
	return r.saveEvent("death_event", e.Date, e.Fact)
}

func (r *Repository) SaveHolidayEvent(e *model.HolidayEvent) error {
	return r.saveEvent("holiday_event", e.Date, e.Fact)
}

func (r *Repository) HistoricalEventIdsBetween(from time.Time, to time.Time) ([]int, error) {
	var ids []int
	results, err := r.db.Query(
		"SELECT id FROM history.historical_event WHERE `date_time` BETWEEN ? AND ?;",
		from,
		to,
	)
	if err != nil {
		return nil, err
	}
	defer func(results *sql.Rows) {
		err := results.Close()
		if err != nil {
			r.logger.Errorf("failed to close results: %v", err)
		}
	}(results)

	for results.Next() {
		var id int
		err = results.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (r *Repository) HistoricalEvent(id int) (model.HistoricalEvent, error) {
	e := model.HistoricalEvent{ID: id}
	if err := r.db.QueryRow(
		"SELECT date_time, fact FROM history.historical_event WHERE id=?",
		id,
	).Scan(&e.Date, &e.Fact); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e, repository.StatusError{Err: err, HTTPCode: http.StatusNotFound}
		}
		return e, err
	}

	return e, nil
}

func (r *Repository) saveEvent(table string, date time.Time, fact string) error {
	_, err := r.db.Exec(
		fmt.Sprintf("INSERT INTO history.%s (date_time, fact) VALUES(?, ?)", table),
		date,
		fact,
	)

	return err
}
