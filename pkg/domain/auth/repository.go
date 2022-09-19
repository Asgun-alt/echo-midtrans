package auth

import (
	"context"
	"echo-midtrans/pkg/domain/users"
)

type DBRepository interface {
	FindByEmail(ctx context.Context, email string) (*users.User, error)
}
