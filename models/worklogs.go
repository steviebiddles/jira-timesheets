package models

type Worklogs struct {
	Total    int       `json:"total"`
	Worklogs []Worklog `json:"worklogs"`
}
