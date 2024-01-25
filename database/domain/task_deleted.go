package domain

import "encoding/json"

type TaskDeleted struct {
	Id        int     `json:"id"`
	TaskId    int     `json:"task_id"`
	DeletedAt []uint8 `json:"deleted_at"`
}

func BodyToTaskDeleted(body []byte) *TaskDeleted {
	if len(body) == 0 {
		return nil
	}

	var task TaskDeleted
	err := json.Unmarshal(body, &task)
	if err != nil {
		return nil
	}

	return &task
}
