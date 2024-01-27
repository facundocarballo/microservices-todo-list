package domain

import "encoding/json"

type MicroserviceSessionSuccessful struct {
	Id             int      `json:"id"`
	MicroserviceId int      `json:"microservice_id"`
	CreatedAt      []uint8  `json:"created_at"`
	EndAt          *[]uint8 `json:"end_at"`
}

func BodyToMicroserviceSessionSuccessful(body []byte) *MicroserviceSessionSuccessful {
	if len(body) == 0 {
		return nil
	}

	var microservice MicroserviceSessionSuccessful
	err := json.Unmarshal(body, &microservice)
	if err != nil {
		return nil
	}

	return &microservice
}
