package transaction

import (
	"echo-midtrans/pkg/domain/campaign"
	"echo-midtrans/pkg/domain/users"

	"github.com/leekchan/accounting"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CampaignID uint   `gorm:"column:campaign_id"`
	UserID     uint   `gorm:"column:user_id"`
	Amount     int    `gorm:"column:amount"`
	Status     string `gorm:"column:status"`
	Code       string `gorm:"column:code"`
	PaymentURL string `gorm:"column:payment_url"`
	User       users.User
	Campaign   campaign.Campaign
}

func (t *Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
