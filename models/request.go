package models

type LoginRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AssignRoleToUserRequestModel struct {
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}

type RegisterRequestModel struct {
	Username        string `json:"username"`
	Name            string `json:"name" validate:"required, min=3,max=20"`
	Surname         string `json:"surname" validate:"required, min=3,max=20"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"required"`
	Password        string `json:"password" validate:"required,min=6"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
}
