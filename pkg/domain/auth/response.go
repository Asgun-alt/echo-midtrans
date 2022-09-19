package auth

import "time"

type Response struct {
	Token     string    `jsong:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}
