package dto

type User struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Email     string `json:"email" validate:"required"`
	DeletedAt string `json:"-"`
}
