package models

import "time"

// HistoricalEvent ...
type HistoricalEvent struct {
	ID   int       `json:"id"`
	Date time.Time `json:"date"`
	Fact string    `json:"fact"`
}

// BirthEvent ...
type BirthEvent struct {
	ID   int       `json:"id"`
	Date time.Time `json:"date"`
	Fact string    `json:"fact"`
}

// DeathEvent ...
type DeathEvent struct {
	ID   int       `json:"id"`
	Date time.Time `json:"date"`
	Fact string    `json:"fact"`
}

// HolidayEvent ...
type HolidayEvent struct {
	ID   int       `json:"id"`
	Date time.Time `json:"date"`
	Fact string    `json:"fact"`
}
