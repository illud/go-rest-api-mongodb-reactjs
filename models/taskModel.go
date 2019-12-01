package models

type Task struct {
	TITLE       string `json:"title"`
	DESCRIPTION string `json:"description"`
	DATE        string `json:"date"`
}

type TaskGet struct {
	ID          string `json:"_id" bson:"_id,omitempty"`
	TITLE       string `json:"title"`
	DESCRIPTION string `json:"description"`
	DATE        string `json:"date"`
}
