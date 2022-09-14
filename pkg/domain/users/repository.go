package users

import "context"

type DBRepository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindByID(ctx context.Context, id uint) (*User, error)
	Create(ctx context.Context, req *User) (*User, error)
	UpdateByID(ctx context.Context, req *User) error
	DeleteByID(ctx context.Context, req *User) error
}
