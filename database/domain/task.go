package domain

import "encoding/json"

type Task struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	CreatedAt   []uint8 `json:"created_at"`
	Description *string `json:"description"`
	CategoryId  *int    `json:"category_id"`
}

func BodyToTask(body []byte) *Task {
	if len(body) == 0 {
		return nil
	}

	var taSK Task
	err := json.Unmarshal(body, &taSK)
	if err != nil {
		return nil
	}

	return &taSK
}
