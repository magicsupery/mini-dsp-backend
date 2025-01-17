package repositories

import (
	"mini-dsp-backend/models"
	"mini-dsp-backend/utils"
)

type UserRepository interface {
	Create(u *models.User) error
	FindByUsername(username string) (*models.User, error)
	FindByID(id int64) (*models.User, error)
}

type userRepo struct{}

func NewUserRepo() UserRepository {
	return &userRepo{}
}

func (r *userRepo) Create(u *models.User) error {
	return utils.DB.Create(u).Error
}

func (r *userRepo) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := utils.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByID(id int64) (*models.User, error) {
	var user models.User
	err := utils.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
