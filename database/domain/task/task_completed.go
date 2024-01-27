package domain

import "encoding/json"

type TaskCompleted struct {
	Id          int     `json:"id"`
	TaskId      int     `json:"task_id"`
	CompletedAt []uint8 `json:"completed_at"`
}

func BodyToTaskCompleted(body []byte) *TaskCompleted {
	if len(body) == 0 {
		return nil
	}

	var task TaskCompleted
	err := json.Unmarshal(body, &task)
	if err != nil {
		return nil
	}

	return &task
}
