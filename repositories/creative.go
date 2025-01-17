package repositories

import (
	"mini-dsp-backend/models"
	"mini-dsp-backend/utils"
)

type CreativeRepository interface {
	Create(cr *models.Creative) error
	FindAll() ([]models.Creative, error)
	FindByID(id int64) (*models.Creative, error)
	Update(cr *models.Creative) error
	Delete(id int64) error
}

type creativeRepo struct{}

func NewCreativeRepo() CreativeRepository {
	return &creativeRepo{}
}

func (r *creativeRepo) Create(cr *models.Creative) error {
	return utils.DB.Create(cr).Error
}

func (r *creativeRepo) FindAll() ([]models.Creative, error) {
	var list []models.Creative
	err := utils.DB.Find(&list).Error
	return list, err
}

func (r *creativeRepo) FindByID(id int64) (*models.Creative, error) {
	var cr models.Creative
	err := utils.DB.First(&cr, id).Error
	if err != nil {
		return nil, err
	}
	return &cr, nil
}

func (r *creativeRepo) Update(cr *models.Creative) error {
	return utils.DB.Save(cr).Error
}

func (r *creativeRepo) Delete(id int64) error {
	return utils.DB.Delete(&models.Creative{}, id).Error
}
