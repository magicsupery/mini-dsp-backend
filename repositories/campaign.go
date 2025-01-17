package repositories

import (
	"mini-dsp-backend/models"
	"mini-dsp-backend/utils"
)

type CampaignRepository interface {
	Create(c *models.Campaign) error
	FindAll() ([]models.Campaign, error)
	FindByID(id int64) (*models.Campaign, error)
	Update(c *models.Campaign) error
	Delete(id int64) error
}

type campaignRepo struct{}

func NewCampaignRepo() CampaignRepository {
	return &campaignRepo{}
}

func (r *campaignRepo) Create(c *models.Campaign) error {
	return utils.DB.Create(c).Error
}

func (r *campaignRepo) FindAll() ([]models.Campaign, error) {
	var list []models.Campaign
	err := utils.DB.Find(&list).Error
	return list, err
}

func (r *campaignRepo) FindByID(id int64) (*models.Campaign, error) {
	var c models.Campaign
	err := utils.DB.First(&c, id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *campaignRepo) Update(c *models.Campaign) error {
	return utils.DB.Save(c).Error
}

func (r *campaignRepo) Delete(id int64) error {
	return utils.DB.Delete(&models.Campaign{}, id).Error
}
