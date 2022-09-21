package http

import (
	"echo-midtrans/cmd/config"
	"echo-midtrans/pkg/domain/common"
	"echo-midtrans/pkg/domain/transaction"
	customMiddleware "echo-midtrans/pkg/middlewares"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TransactionHTTPHandler struct {
	common.BaseHTTPHandler
	usecase transaction.Usecase
}

func NewTransactionHTTPHandler(appGroup *echo.Group, transactionService transaction.Usecase) {
	handler := &TransactionHTTPHandler{usecase: transactionService}

	jwtConfig := customMiddleware.NewJWTMiddlewareConfig()
	transactionGroup := appGroup.Group("/transaction")
	transactionGroup.GET("/campaign/:id", handler.FindCampaignTransactions)
	transactionGroup.GET("/user/:id", handler.FindUserTransactions)
	transactionGroup.GET("/notification", handler.GetNotification)
	transactionGroup.POST("", handler.CreateTransaction, middleware.JWTWithConfig(jwtConfig))
}

func (h *TransactionHTTPHandler) FindCampaignTransactions(ctx echo.Context) error {
	var (
		id  int
		err error
	)

	idStr := ctx.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.InvalidCampaignID, http.StatusBadRequest)
	}

	campaigns, err := h.usecase.GetTransactionsByCampaignID(ctx.Request().Context(), uint(id))
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.ErrRecordNotFound, http.StatusBadRequest)
		}
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, campaigns, nil, http.StatusCreated)
}

func (h *TransactionHTTPHandler) FindUserTransactions(ctx echo.Context) error {
	var (
		id  int
		err error
	)
	idStr := ctx.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.InvalidUserID, http.StatusBadRequest)
	}

	campaigns, err := h.usecase.GetTransactionsByCampaignID(ctx.Request().Context(), uint(id))
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.ErrRecordNotFound, http.StatusBadRequest)
		}
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, campaigns, nil, http.StatusCreated)
}

func (h *TransactionHTTPHandler) CreateTransaction(ctx echo.Context) error {
	var (
		request transaction.CreateTransactionRequest
		err     error
	)
	valid := ctx.Get("validator").(*config.CustomValidator)
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*common.JWTCustomClaims)
	userID := claims.UserID

	request.UserID = userID

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

	newTransaction, err := h.usecase.CreateTransaction(ctx.Request().Context(), &request)
	if err != nil {
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, newTransaction, nil, http.StatusCreated)
}

func (h *TransactionHTTPHandler) GetNotification(ctx echo.Context) error {
	var (
		request transaction.TransactionNotificationRequest
		err     error
	)

	err = ctx.Bind(&request)
	if err != nil {
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusUnprocessableEntity)
	}

	err = h.usecase.ProcessPayment(ctx.Request().Context(), &request)
	if err != nil {
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusUnprocessableEntity)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, nil, nil, http.StatusOK)
}
