package repositories

import (
	"mini-dsp-backend/models"
	"mini-dsp-backend/utils"
)

type AdvertiserRepository interface {
	Create(ad *models.Advertiser) error
	FindAll() ([]models.Advertiser, error)
	FindByID(id int64) (*models.Advertiser, error)
	Update(ad *models.Advertiser) error
	Delete(id int64) error
}

type advertiserRepo struct{}

func NewAdvertiserRepo() AdvertiserRepository {
	return &advertiserRepo{}
}

func (r *advertiserRepo) Create(ad *models.Advertiser) error {
	return utils.DB.Create(ad).Error
}

func (r *advertiserRepo) FindAll() ([]models.Advertiser, error) {
	var list []models.Advertiser
	err := utils.DB.Find(&list).Error
	return list, err
}

func (r *advertiserRepo) FindByID(id int64) (*models.Advertiser, error) {
	var ad models.Advertiser
	err := utils.DB.First(&ad, id).Error
	if err != nil {
		return nil, err
	}
	return &ad, nil
}

func (r *advertiserRepo) Update(ad *models.Advertiser) error {
	return utils.DB.Save(ad).Error
}

func (r *advertiserRepo) Delete(id int64) error {
	return utils.DB.Delete(&models.Advertiser{}, id).Error
}
