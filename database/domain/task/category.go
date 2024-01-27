package domain

import "encoding/json"

type Category struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	CreatedAt   []uint8 `json:"created_at"`
	Description *string `json:"description"`
	PhotoUrl    *string `json:"photo_url"`
	ParentId    *int    `json:"parent_id"`
}

func BodyToCategory(body []byte) *Category {
	if len(body) == 0 {
		return nil
	}

	var category Category
	err := json.Unmarshal(body, &category)
	if err != nil {
		return nil
	}

	return &category
}
