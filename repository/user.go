package repository

import (
	"example.com/awesome/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	FindAll() []model.User
	FindByID(id uint) (*model.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) FindAll() []model.User {
	var users []model.User
	if err := repo.db.Find(&users).Error; err != nil {
		panic(err)
	}
	return users
}

func (repo *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
