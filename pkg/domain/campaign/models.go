package campaign

import (
	"github.com/leekchan/accounting"
	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	CampaignName  string `gorm:"column:campaign_name"`
	Description   string `gorm:"column:description"`
	Perks         string `gorm:"column:perks"`
	BackerCount   int    `gorm:"column:backer_count"`
	GoalAmount    int    `gorm:"column:goal_amount"`
	CurrentAmount int    `gorm:"column:current_amount"`
	Slug          string `gorm:"column:slug"`
}

func (c *Campaign) ToResponse() *Response {
	return &Response{
		CampaignName:  c.CampaignName,
		Description:   c.Description,
		Perks:         c.Perks,
		BackerCount:   c.BackerCount,
		GoalAmount:    c.GoalAmount,
		CurrentAmount: c.CurrentAmount,
		Slug:          c.Slug,
	}
}

func ToMultipleResponse(req []Campaign) (output []Response) {
	for idx := range req {
		output = append(output, Response{
			ID:            req[idx].ID,
			CampaignName:  req[idx].CampaignName,
			Description:   req[idx].Description,
			Perks:         req[idx].Perks,
			BackerCount:   req[idx].BackerCount,
			GoalAmount:    req[idx].GoalAmount,
			CurrentAmount: req[idx].CurrentAmount,
		})
	}
	return output
}

type CampaignImage struct {
	gorm.Model
	CampaignID uint   `gorm:"column:campaign_id"`
	FileName   string `gorm:"column:file_name"`
}

func (c Campaign) GoalAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(c.GoalAmount)
}

func (c Campaign) CurrentAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(c.CurrentAmount)
}
