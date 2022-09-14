package usecase

import "echo-midtrans/pkg/domain/users"

type UsersUseCase struct {
	DBRepo users.DBRepository
}

func NewUsersUseCase(dbrepo users.DBRepository) *UsersUseCase {
	return &UsersUseCase{
		DBRepo: dbrepo,
	}
}
