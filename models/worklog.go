package models

type Worklog struct {
	Id               string  `json:"id,omitempty"`
	Self             string  `json:"self,omitempty"`
	Author           Author  `json:"author,omitempty"`
	Comment          Comment `json:"comment,omitempty"`
	Started          string  `json:"started"`
	TimeSpent        string  `json:"timeSpent,omitempty"`
	TimeSpentSeconds int     `json:"timeSpentSeconds,omitempty"`
}
