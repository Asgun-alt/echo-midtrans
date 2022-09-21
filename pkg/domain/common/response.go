package common

import "github.com/golang-jwt/jwt"

type JWTCustomClaims struct {
	UserID uint
	jwt.StandardClaims
}
