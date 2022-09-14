package repository

import "gorm.io/gorm"

type UsersDBRepository struct {
	DB *gorm.DB
}

func NewUsersDBRepository(db *gorm.DB) *UsersDBRepository {
	return &UsersDBRepository{
		DB: db,
	}
}
