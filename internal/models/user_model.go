package models

import "time"

type CreateUser struct {
	Username  string    `json:"username" validate:"required,alphanumspecial"`
	Password  string    `json:"password" validate:"required,notallowchar"`
	Email     string    `json:"email" validate:"required,email"`
	FullName  string    `json:"full_name" validate:"required,notallowchar"`
	BirthDate time.Time `json:"birth_date" validate:"required"`
}

type LoginUser struct {
	Password string `json:"password" validate:"required,alphanum"`
	Email    string `json:"email" validate:"required,email"`
}

type UpdateUser struct {
	Username  string    `json:"username" validate:"required,alphanumspecial"`
	FullName  string    `json:"full_name" validate:"required,notallowchar"`
	BirthDate time.Time `json:"birth_date" validate:"required"`
}

type ChangePassword struct {
	OldPassword string `json:"old_password" validate:"required,notallowchar"`
	NewPassword string `json:"new_password" validate:"required,notallowchar"`
}
