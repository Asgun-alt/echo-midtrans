package campaign

import (
	"gorm.io/gorm"
)

type CampaignDetailsRequest struct {
	ID int `uri:"id"`
}

type CreateCampaignRequest struct {
	CampaignName string `json:"campaign_name" validate:"max=150"`
	Description  string `json:"description" validate:"max=350"`
	GoalAmount   int    `json:"goal_amount"`
	Perks        string `json:"perks" validate:"max=100"`
}

func (data *CreateCampaignRequest) ToCampaignDomain() *Campaign {
	return &Campaign{
		Model:        gorm.Model{},
		CampaignName: data.CampaignName,
		Description:  data.Description,
		GoalAmount:   data.GoalAmount,
		Perks:        data.Perks,
	}
}

type UpdateCampaignRequest struct {
	ID           uint   `json:"id"`
	CampaignName string `json:"campaign_name" validate:"max=150"`
	Description  string `json:"description" validate:"max=350"`
	GoalAmount   int    `json:"goal_amount"`
	Perks        string `json:"perks" validate:"max=100"`
}

func (data *UpdateCampaignRequest) ToCampaignDomain() *Campaign {
	return &Campaign{
		Model: gorm.Model{
			ID: uint(data.ID),
		},
		CampaignName: data.CampaignName,
		Description:  data.Description,
		GoalAmount:   data.GoalAmount,
		Perks:        data.Perks,
	}
}

type CreateCampaignImageRequest struct {
	CampaignID uint `form:"campaign_id"`
	IsPrimary  bool `form:"is_primary"`
}

func (data *CreateCampaignImageRequest) ToCampaignDomain() *CampaignImage {
	return &CampaignImage{
		CampaignID: data.CampaignID,
		IsPrimary:  data.IsPrimary,
	}
}
