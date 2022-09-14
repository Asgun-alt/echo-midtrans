package users

type Response struct {
	ID       uint
	Username string `json:"user_name"`
	Email    string `json:"email"`
}
