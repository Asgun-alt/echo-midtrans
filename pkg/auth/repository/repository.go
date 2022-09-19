package repository

import (
	"context"
	"echo-midtrans/pkg/domain/users"
	"fmt"

	"gorm.io/gorm"
)

type AuthDBRepository struct {
	DB *gorm.DB
}

func NewAuthDBRepository(db *gorm.DB) *AuthDBRepository {
	return &AuthDBRepository{
		DB: db,
	}
}

func (r *AuthDBRepository) FindByEmail(ctx context.Context, email string) (*users.User, error) {
	var res users.User

	err := r.DB.WithContext(ctx).First(&res, "email = ?", email).Error
	if err != nil {
		return nil, fmt.Errorf("AuthDBRepository.FindByEmail: %w", err)
	}

	return &res, nil
}
