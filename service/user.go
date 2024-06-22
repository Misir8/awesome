package service

import "example.com/awesome/model"

func GetUser() model.User {
	return model.User{
		Name: "John Doe",
	}
}
