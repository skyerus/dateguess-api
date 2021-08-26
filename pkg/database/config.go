package database

import (
	"net/url"
)

// Config holds the config for the database connection
type Config struct {
	MySQLUser     string
	MySQLPassword string
	MySQLHost     string
	MySQLOptions  string
	MySQLTimeZone string
}

// ConnectionString returns the database connection string
func (c *Config) ConnectionString() string {
	return c.MySQLUser +
		":" +
		c.MySQLPassword +
		"@tcp(" +
		c.MySQLHost +
		")/?" +
		c.MySQLOptions +
		"&time_zone=" +
		url.QueryEscape(c.MySQLTimeZone)
}
