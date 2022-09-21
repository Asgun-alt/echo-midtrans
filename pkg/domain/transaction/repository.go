package transaction

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]Transaction, error)
	FindByCampaignID(ctx context.Context, campaignID uint) ([]Transaction, error)
	FindByUserID(ctx context.Context, userId uint) ([]Transaction, error)
	Create(ctx context.Context, transaction *Transaction) (*Transaction, error)
	Update(ctx context.Context, transaction *Transaction) error
}
