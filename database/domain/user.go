package domain

import "encoding/json"

type User struct {
	Id        int     `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	CreatedAt []uint8 `json:"created_at"`
	PhotoUrl  *string `json:"photo_url"`
}

func BodyToUser(body []byte) *User {
	if len(body) == 0 {
		return nil
	}

	var user User
	err := json.Unmarshal(body, &user)
	if err != nil {
		return nil
	}

	return &user
}
