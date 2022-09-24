package token

import "time"

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}
type TokenPayload struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
}
