package common

import "github.com/labstack/echo/v4"

type BaseHTTPHandler struct{}

func (h *BaseHTTPHandler) ResponseJSON(ctx echo.Context, message string, data, errors interface{}, httpCode int) error {
	return ctx.JSON(httpCode, Response{
		Message: message,
		Data:    data,
		Errors:  errors,
		Code:    httpCode,
	})
}
