package middlewares

import (
	"echo-midtrans/pkg/domain/common"

	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func NewJWTMiddlewareConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &common.JWTCustomClaims{},
		SigningKey: []byte(viper.GetString("jwt.SecretKey")),
	}
}
