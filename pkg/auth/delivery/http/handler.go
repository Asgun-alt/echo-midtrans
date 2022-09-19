package http

import (
	"echo-midtrans/cmd/config"
	"echo-midtrans/pkg/domain/auth"
	"echo-midtrans/pkg/domain/common"
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AuthHTTPHandler struct {
	common.BaseHTTPHandler
	Usecase auth.UseCase
}

func NewAuthHTTPHandler(appGroup *echo.Group, uc auth.UseCase) {
	handler := &AuthHTTPHandler{
		Usecase: uc,
	}

	usersGroup := appGroup.Group("/auth")
	usersGroup.POST("/login", handler.ValidateUser)
}

func (h *AuthHTTPHandler) ValidateUser(ctx echo.Context) error {
	var (
		request  auth.ValidateUserRequest
		response *auth.Response
		err      error
	)
	valid := ctx.Get("validator").(*config.CustomValidator)

	err = ctx.Bind(&request)
	if err != nil {
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusUnprocessableEntity)
	}

	err = ctx.Validate(&request)
	if err != nil {
		if valErr, ok := err.(validator.ValidationErrors); ok {
			return h.ResponseJSON(ctx, common.ValidationError, nil, valErr.Translate(valid.Translator), http.StatusBadRequest)
		}
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusNotFound)
	}

	response, err = h.Usecase.ValidateUser(ctx.Request().Context(), &request)
	if err != nil {
		if errors.Is(err, common.ErrPasswordNotMatch) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.PasswordNotMatch, http.StatusBadRequest)
		}
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, response, nil, http.StatusOK)
}
