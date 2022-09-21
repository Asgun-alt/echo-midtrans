package payment

import (
	"context"
	"echo-midtrans/pkg/domain/transaction"
	"echo-midtrans/pkg/domain/users"
)

type UseCase interface {
	GetPaymentURL(ctx context.Context, payment *transaction.Transaction, user users.User) (string, error)
}
