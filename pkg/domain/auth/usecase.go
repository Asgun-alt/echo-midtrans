package auth

import "context"

type UseCase interface {
	ValidateUser(Ctx context.Context, req *ValidateUserRequest) (*Response, error)
}
