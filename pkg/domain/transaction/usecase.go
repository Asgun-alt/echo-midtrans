package transaction

import "context"

type Usecase interface {
	GetTransactionsByCampaignID(ctx context.Context, campaignID uint) ([]Transaction, error)
	GetTransactionsByUserID(ctx context.Context, userID uint) ([]Transaction, error)
	CreateTransaction(ctx context.Context, req *CreateTransactionRequest) (*Transaction, error)
	ProcessPayment(ctx context.Context, req *TransactionNotificationRequest) error
}
