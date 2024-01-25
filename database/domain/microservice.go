package domain

import "encoding/json"

type Microservice struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	LanguageId  int     `json:"language_id"`
	PublicKey   string  `json:"public_key"`
	Description *string `json:"description"`
}

func BodyToMicroservice(body []byte) *Microservice {
	if len(body) == 0 {
		return nil
	}

	var microservice Microservice
	err := json.Unmarshal(body, &microservice)
	if err != nil {
		return nil
	}

	return &microservice
}
