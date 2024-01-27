package domain

import "encoding/json"

type MicroserviceSessionRejected struct {
	Id             int     `json:"id"`
	MicroserviceId int     `json:"microservice_id"`
	CreatedAt      []uint8 `json:"created_at"`
}

func BodyToMicroserviceSessionRejected(body []byte) *MicroserviceSessionRejected {
	if len(body) == 0 {
		return nil
	}

	var microservice MicroserviceSessionRejected
	err := json.Unmarshal(body, &microservice)
	if err != nil {
		return nil
	}

	return &microservice
}
