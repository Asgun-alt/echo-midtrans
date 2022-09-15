package http

import (
	"echo-midtrans/cmd/config"
	"echo-midtrans/pkg/domain/campaign"
	"echo-midtrans/pkg/domain/common"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CampaignHTTPHandler struct {
	common.BaseHTTPHandler
	usecase campaign.UseCase
}

func NewCampaignHTTPHandler(appGroup echo.Group, campaignService campaign.UseCase) {
	handler := &CampaignHTTPHandler{usecase: campaignService}

	campaignGroup := appGroup.Group("/campaign")
	campaignGroup.GET("", handler.FindAll)
	campaignGroup.GET("/details", handler.FindCampaignDetails)
	campaignGroup.POST("", handler.AddCampaign)
	campaignGroup.PUT("/:id", handler.UpdateCampaign)
	campaignGroup.DELETE("/:id", handler.DeleteCampaign)

	campaignGroup.POST("/upload", handler.UploadCampaignImage)
}

func (h *CampaignHTTPHandler) FindAll(ctx echo.Context) error {
	res, err := h.usecase.GetCampaigns(ctx.Request().Context())
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.RecordNotFound, http.StatusNotFound)
		}
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}
	return h.ResponseJSON(ctx, common.DataSuccess, campaign.ToMultipleResponse(res), nil, http.StatusOK)
}

func (h *CampaignHTTPHandler) FindCampaignDetails(ctx echo.Context) error { return nil }
func (h *CampaignHTTPHandler) UploadCampaignImage(ctx echo.Context) error { return nil }

func (h *CampaignHTTPHandler) AddCampaign(ctx echo.Context) error {
	var (
		request  campaign.CreateCampaignRequest
		campaign *campaign.Campaign
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

	campaign, err = h.usecase.CreateCampaign(ctx.Request().Context(), request.ToCampaignDomain())
	if err != nil {
		if errors.Is(err, common.ErrCampaignAlreadyCreated) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.CampaignAlreadyCreated, http.StatusBadRequest)
		}
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}
	return h.ResponseJSON(ctx, common.DataSuccess, campaign.ToResponse(), nil, http.StatusCreated)

}

func (h *CampaignHTTPHandler) UpdateCampaign(ctx echo.Context) error {
	var (
		request campaign.UpdateCampaignRequest
		id      int
		err     error
	)

	idStr := ctx.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.InvalidUserID, http.StatusBadRequest)
	}

	campaignCustomValidator := ctx.Get("validator").(*config.CustomValidator)
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
	if request.CampaignName != "" {
		err = (campaignCustomValidator).Validator.Var(&request.CampaignName, "max=150")
		if err != nil {
			valErr := err.(validator.ValidationErrors)
			return h.ResponseJSON(ctx, common.ValidationError, nil, valErr.Translate(campaignCustomValidator.Translator), http.StatusBadRequest)
		}
	}

	err = h.usecase.UpdateCampaign(ctx.Request().Context(), request.ToCampaignDomain())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.RecordNotFound, http.StatusNotFound)
		}

		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, nil, nil, http.StatusOK)
}

func (h *CampaignHTTPHandler) DeleteCampaign(ctx echo.Context) error {
	var (
		id  int
		err error
	)
	idStr := ctx.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.InvalidUserID, http.StatusBadRequest)
	}

	err = h.usecase.DeleteCampaign(ctx.Request().Context(), &campaign.Campaign{
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
