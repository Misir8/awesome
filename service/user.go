package service

import (
	"example.com/awesome/database"
	"example.com/awesome/model"
	"example.com/awesome/repository"
)

type UserService struct {
	UserRepo repository.IUserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo: repository.NewUserRepository(database.DbManager),
	}
}

func (s *UserService) GetUser(id uint) (*model.User, error) {
	return s.UserRepo.FindByID(id)
}
