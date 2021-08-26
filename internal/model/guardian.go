package model

import "time"

// GuardianContent ...
type GuardianContent struct {
	Response struct {
		Status      string           `json:"status"`
		UserTier    string           `json:"userTier"`
		Total       int              `json:"total"`
		StartIndex  int              `json:"startIndex"`
		PageSize    int              `json:"pageSize"`
		CurrentPage int              `json:"currentPage"`
		Pages       int              `json:"pages"`
		OrderBy     string           `json:"orderBy"`
		Results     []GuardianResult `json:"results"`
	} `json:"response"`
}

// GuardianResult ...
type GuardianResult struct {
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

type SearchParams struct {
	PageSize int
	Page     int
	Section  string
	From     time.Time
	To       time.Time
}
