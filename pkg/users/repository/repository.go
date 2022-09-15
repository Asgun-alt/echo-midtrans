package repository

import (
	"context"
	"echo-midtrans/pkg/domain/common"
	"echo-midtrans/pkg/domain/users"
	"errors"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type UsersDBRepository struct {
	DB *gorm.DB
}

func NewUsersDBRepository(db *gorm.DB) *UsersDBRepository {
	return &UsersDBRepository{
		DB: db,
	}
}

func (r *UsersDBRepository) FindAll(ctx context.Context) ([]users.User, error) {
	var (
		res []users.User
	)

	err := r.DB.WithContext(ctx).Find(&res).Order("created_at DESC").Error
	if err != nil {
		return nil, fmt.Errorf("UsersDBRepository.FindAll: %w", err)
	}
	if res == nil {
		return nil, common.ErrRecordNotFound
	}
	return res, nil
}

func (r *UsersDBRepository) FindByID(ctx context.Context, ID uint) (*users.User, error) {
	var (
		res users.User
	)

	err := r.DB.WithContext(ctx).First(&res, "id = ?", ID).Error
	if err != nil {
		return nil, fmt.Errorf("UsersDBRepository.FindByID: %w", err)
	}
	return &res, nil
}

func (r *UsersDBRepository) Create(ctx context.Context, req *users.User) (*users.User, error) {
	var (
		res users.User
	)

	err := r.DB.WithContext(ctx).First(&res, "username = ?", req.Username).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("UsersDBRepository.Create: %w", err)
	}

	// DeepEqual is used to check two interfaces are equal or not
	if !reflect.DeepEqual(res, users.User{}) {
		return nil, common.ErrUserAlreadyCreated
	}

	err = r.DB.WithContext(ctx).Save(req).Error
	if err != nil {
		return nil, fmt.Errorf("UsersDBRepository.Create: %w", err)
	}
	return req, nil
}

func (r *UsersDBRepository) UpdateByID(ctx context.Context, req *users.User) error {
	err := r.DB.WithContext(ctx).Updates(req).Error
	if err != nil {
		return fmt.Errorf("UsersDBRepository.UpdateByID: %w", err)
	}
	return nil
}

func (r *UsersDBRepository) DeleteByID(ctx context.Context, req *users.User) error {
	err := r.DB.WithContext(ctx).Delete(req).Error
	if err != nil {
		return fmt.Errorf("UsersDBRepository.DeleteByID: %w", err)
	}

	return nil
}
