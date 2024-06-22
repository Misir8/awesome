package dto

import "example.com/awesome/model"

type ResponseUserDto struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func UserDTOFromModel(user *model.User) *ResponseUserDto {
	if user == nil {
		return nil
	}
	return &ResponseUserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}
}

type RegisterUserDto struct {
	Phone    string `json:"phone" validate:"required"`
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
