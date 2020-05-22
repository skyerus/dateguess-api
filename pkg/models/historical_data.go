package models

import "time"

// HistoricalEvent ...
type HistoricalEvent struct {
	Date time.Time
	Fact string
}

// BirthEvent ...
type BirthEvent struct {
	Date time.Time
	Fact string
}

// DeathEvent ...
type DeathEvent struct {
	Date time.Time
	Fact string
}

// HolidayEvent ...
type HolidayEvent struct {
	Date time.Time
	Fact string
}
