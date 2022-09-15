package users

import "gorm.io/gorm"

type AddUserRequest struct {
	Username  string `json:"user_name" validate:"required,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Password1 string `json:"password_1" validate:"required,eqfield=Password2"`
	Password2 string `json:"password_2" validate:"required"`
}

func (data *AddUserRequest) ToUserDomain() *User {
	return &User{
		Model:    gorm.Model{},
		Username: data.Username,
		Password: data.Password1,
		Email:    data.Email,
	}
}

type UpdateUserRequest struct {
	ID               uint   `json:"id" validate:"required"`
	Username         string `json:"user_name"`
	Email            string `json:"email"`
	ExistingPassword string `json:"existing_password"`
	Password1        string `json:"password_1"`
	Password2        string `json:"password_2"`
}

func (data *UpdateUserRequest) ToUserDomain() *User {
	return &User{
		Model: gorm.Model{
			ID: data.ID,
		},
		Username: data.Username,
		Password: data.Password1,
		Email:    data.Email,
	}
}
