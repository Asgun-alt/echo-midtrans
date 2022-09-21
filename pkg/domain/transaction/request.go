package transaction

import "echo-midtrans/pkg/domain/users"

type CampaignTransactionRequest struct {
	ID   uint `uri:"id" validate:"required"`
	User users.User
}

type CreateTransactionRequest struct {
	Amount     int  `json:"amount" validate:"required"`
	CampaignID uint `json:"campaign_id" validate:"required"`
	UserID     uint
}

type TransactionNotificationRequest struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           uint   `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
