package model

import "time"

// Article ...
type Article struct {
	ID                 string    `json:"id"`
	Type               string    `json:"type"`
	SectionID          string    `json:"sectionId"`
	SectionName        string    `json:"sectionName"`
	WebPublicationDate time.Time `json:"webPublicationDate"`
	WebTitle           string    `json:"webTitle"`
	WebURL             string    `json:"webUrl"`
	APIURL             string    `json:"apiUrl"`
	Fields             struct {
		BodyText string `json:"bodyText"`
	} `json:"fields"`
	IsHosted   bool   `json:"isHosted"`
	PillarID   string `json:"pillarId"`
	PillarName string `json:"pillarName"`
}
