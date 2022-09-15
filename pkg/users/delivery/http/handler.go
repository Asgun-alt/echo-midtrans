package http

import (
	"echo-midtrans/cmd/config"
	"echo-midtrans/pkg/domain/common"
	"echo-midtrans/pkg/domain/users"
	"echo-midtrans/pkg/helpers"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UsersHTTPHandler struct {
	common.BaseHTTPHandler
	usecase users.UseCase
}

func NewUsersHTTPHandler(appGroup *echo.Group, uc users.UseCase) {
	handler := &UsersHTTPHandler{
		usecase: uc,
	}

	usersGroup := appGroup.Group("/users")
	usersGroup.GET("", handler.FindAll)
	usersGroup.POST("", handler.AddUser)
	usersGroup.PUT("/:id", handler.UpdateUser)
	usersGroup.DELETE("/:id", handler.DeleteUser)
}

func (h *UsersHTTPHandler) FindAll(ctx echo.Context) error {
	res, err := h.usecase.FindAll(ctx.Request().Context())
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.RecordNotFound, http.StatusNotFound)
		}
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}
	return h.ResponseJSON(ctx, common.DataSuccess, users.ToMultipleResponse(res), nil, http.StatusOK)
}

func (h *UsersHTTPHandler) AddUser(ctx echo.Context) error {
	var (
		request users.AddUserRequest
		user    *users.User
		err     error
	)
	valid := ctx.Get("validator").(*config.CustomValidator)

	err = ctx.Bind(&request)
	if err != nil {
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusUnprocessableEntity)
	}

	err = ctx.Validate(&request)
	if err != nil {
		fmt.Println("debug1")
		if valErr, ok := err.(validator.ValidationErrors); ok {
			fmt.Println("debug2")
			return h.ResponseJSON(ctx, common.ValidationError, nil, valErr.Translate(valid.Translator), http.StatusBadRequest)
		}
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusNotFound)
	}

	isValidPassword := helpers.IsPasswordOK(request.Password1)
	if !isValidPassword {
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.InvalidPassword, http.StatusBadRequest)
	}

	user, err = h.usecase.CreateUser(ctx.Request().Context(), request.ToUserDomain())
	if err != nil {
		if errors.Is(err, common.ErrUserAlreadyCreated) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.UserAlreadyCreated, http.StatusBadRequest)
		}
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, user.ToResponse(), nil, http.StatusCreated)
}

func (h *UsersHTTPHandler) UpdateUser(ctx echo.Context) error {
	var (
		request users.UpdateUserRequest
		id      int
		err     error
	)

	idStr := ctx.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.InvalidUserID, http.StatusBadRequest)
	}

	userCustomValidator := ctx.Get("validator").(*config.CustomValidator)
	err = ctx.Bind(&request)
	if err != nil {
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusUnprocessableEntity)
	}

	err = ctx.Validate(&request)
	if err != nil {
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusBadRequest)
	}
	if uint(id) != request.ID {
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.InvalidUserID, http.StatusBadRequest)
	}

	if request.ExistingPassword != "" {
		if request.Password1 == "" {
			return h.ResponseJSON(ctx, common.ValidationError, nil, common.PasswordEmpty, http.StatusBadRequest)
		}

		if request.Password1 != request.Password2 {
			return h.ResponseJSON(ctx, common.ValidationError, nil, common.PasswordNotSame, http.StatusBadRequest)
		}

		isValidPassword := helpers.IsPasswordOK(request.Password1)
		if !isValidPassword {
			return h.ResponseJSON(ctx, common.ValidationError, nil, common.InvalidPassword, http.StatusBadRequest)
		}
	} else {
		// To prevent updating the password when the existing password is not filled
		request.Password1, request.Password2 = "", ""
	}
	if request.Username != "" {
		err = (userCustomValidator).Validator.Var(&request.Username, "max=50")
		if err != nil {
			valErr := err.(validator.ValidationErrors)
			return h.ResponseJSON(ctx, common.ValidationError, nil, valErr.Translate(userCustomValidator.Translator), http.StatusBadRequest)
		}
	}
	if request.Email != "" {
		err = (userCustomValidator).Validator.Var(&request.Username, "max=50")
		if err != nil {
			valErr := err.(validator.ValidationErrors)
			return h.ResponseJSON(ctx, common.ValidationError, nil, valErr.Translate(userCustomValidator.Translator), http.StatusBadRequest)
		}
	}

	err = h.usecase.UpdateUser(ctx.Request().Context(), request.ExistingPassword, request.ToUserDomain())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.RecordNotFound, http.StatusNotFound)
		}

		if errors.Is(err, common.ErrPasswordNotMatch) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.PasswordNotMatch, http.StatusBadRequest)
		}

		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, nil, nil, http.StatusOK)
}

func (h *UsersHTTPHandler) DeleteUser(ctx echo.Context) error {
	var (
		id  int
		err error
	)
	idStr := ctx.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.InvalidUserID, http.StatusBadRequest)
	}

	err = h.usecase.DeleteUser(ctx.Request().Context(), &users.User{
		Model: gorm.Model{
			ID: uint(id),
		},
	})
	if err != nil {
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, nil, nil, http.StatusOK)
}
