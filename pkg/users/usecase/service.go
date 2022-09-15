package usecase

import (
	"context"
	"echo-midtrans/pkg/domain/common"
	"echo-midtrans/pkg/domain/users"
	"echo-midtrans/pkg/helpers"
)

type UsersUseCase struct {
	DBRepo users.DBRepository
}

func NewUsersUseCase(dbrepo users.DBRepository) *UsersUseCase {
	return &UsersUseCase{
		DBRepo: dbrepo,
	}
}

func (uc *UsersUseCase) FindAll(ctx context.Context) ([]users.User, error) {
	users, err := uc.DBRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *UsersUseCase) CreateUser(ctx context.Context, req *users.User) (*users.User, error) {
	var (
		err error
	)

	req.Password, err = helpers.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	req, err = uc.DBRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (uc *UsersUseCase) UpdateUser(ctx context.Context, existingPassword string, req *users.User) error {
	user, err := uc.DBRepo.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if existingPassword != "" {
		if !helpers.CheckPasswordHash(existingPassword, user.Password) {
			return common.ErrPasswordNotMatch
		}

		req.Password, err = helpers.HashPassword(req.Username)
		if err != nil {
			return err
		}
	}

	err = uc.DBRepo.UpdateByID(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UsersUseCase) DeleteUser(ctx context.Context, req *users.User) error {
	err := uc.DBRepo.DeleteByID(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
