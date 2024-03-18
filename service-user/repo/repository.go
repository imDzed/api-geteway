package repository

import (
	"errors"
	"service-user/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (userRepo *userRepositoryImpl) Create(user *model.User) error {
	result := userRepo.db.Create(user)
	if result.Error != nil {
		return errors.New("error creating user")
	}
	return nil
}

func (userRepo *userRepositoryImpl) FindUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	result := userRepo.db.Where("email = ?", email).First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("error checking email existence")
	}
	return user, nil
}
