package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
}

func (u *User) ToResponse() *Response {
	return &Response{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}

func ToMultipleResponse(req []User) (output []Response) {
	for idx := range req {
		output = append(output, Response{
			ID:       req[idx].ID,
			Username: req[idx].Username,
			Email:    req[idx].Email,
		})
	}
	return output
}
