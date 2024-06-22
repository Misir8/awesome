package service

import (
	"example.com/awesome/database"
	"example.com/awesome/dto"
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

func (s *UserService) GetUser(id uint) (*dto.ResponseUserDto, error) {
	user, err := s.UserRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return dto.UserDTOFromModel(user), nil
}

func (s *UserService) GetUsers() ([]*dto.ResponseUserDto, error) {
	users, err := s.UserRepo.FindAll()
	if err != nil {
		return nil, err
	}
	var userDTOs []*dto.ResponseUserDto
	for _, user := range users {
		userDTOs = append(userDTOs, dto.UserDTOFromModel(&user))
	}

	return userDTOs, nil
}
