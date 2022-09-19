package common

import "github.com/golang-jwt/jwt"

type JWTCustomClaims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}
