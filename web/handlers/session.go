package handlers

import (
	"echo-midtrans/pkg/domain/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SessionHandler struct {
	service users.UseCase
}

func NewSessionHandler(userService users.UseCase) *SessionHandler {
	return &SessionHandler{service: userService}
}

func (h *SessionHandler) New(ctx echo.Context) {
	ctx.HTML(http.StatusOK, "session_new.html")
}

// func (h *SessionHandler) Create(ctx echo.Context) {
// 	var request auth.ValidateUserRequest

// }
