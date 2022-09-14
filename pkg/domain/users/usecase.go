package users

import "context"

type UseCase interface {
	FindAll(ctx context.Context) ([]User, error)
	CreateUser(ctx context.Context, req *User) (*User, error)
	UpdateUser(ctx context.Context, existingPassword string, req *User) error
	DeleteUser(ctx context.Context, req *User) error
}
