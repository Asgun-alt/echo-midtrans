package auth

import (
	"time"
)

type Response struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}
