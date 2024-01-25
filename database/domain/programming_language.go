package domain

import "encoding/json"

type ProgrammingLanguage struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func BodyToProgrammingLanguage(body []byte) *ProgrammingLanguage {
	if len(body) == 0 {
		return nil
	}

	var pl ProgrammingLanguage
	err := json.Unmarshal(body, &pl)
	if err != nil {
		return nil
	}

	return &pl
}
