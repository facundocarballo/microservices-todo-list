package domain

import "encoding/json"

type Task struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	CreatedAt   []uint8 `json:"created_at"`
	CategoryId  int     `json:"category_id"`
	UserId      int     `json:"user_id"`
	Description *string `json:"description"`
}

func BodyToTask(body []byte) *Task {
	if len(body) == 0 {
		return nil
	}

	var task Task
	err := json.Unmarshal(body, &task)
	if err != nil {
		return nil
	}

	return &task
}
