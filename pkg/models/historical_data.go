package models

import "time"

// HistoricalEvent ...
type HistoricalEvent struct {
	ID   int
	Date time.Time
	Fact string
}

// BirthEvent ...
type BirthEvent struct {
	ID   int
	Date time.Time
	Fact string
}

// DeathEvent ...
type DeathEvent struct {
	ID   int
	Date time.Time
	Fact string
}

// HolidayEvent ...
type HolidayEvent struct {
	ID   int
	Date time.Time
	Fact string
}
