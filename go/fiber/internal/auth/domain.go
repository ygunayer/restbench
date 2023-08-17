package auth

import "github.com/ygunayer/restbench/internal/user"

type RegisterUserRequest struct {
	Name                 string `json:"name" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" validate:"eqfield=Password"`
}

func (r *RegisterUserRequest) ToCommand() user.RegisterUserCommand {
	return user.RegisterUserCommand{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}
