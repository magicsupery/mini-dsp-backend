package services

import (
	"errors"
	"time"

	"mini-dsp-backend/models"
	"mini-dsp-backend/repositories"
)

// CampaignService 定义了活动管理的业务接口
type CampaignService interface {
	CreateCampaign(advertiserID int64, name string, budget, bidAmount float64, bidType string) (*models.Campaign, error)
	ListCampaigns() ([]models.Campaign, error)
	GetCampaignByID(id int64) (*models.Campaign, error)
	UpdateCampaign(campaign *models.Campaign) (*models.Campaign, error)
	DeleteCampaign(id int64) error
}

// campaignService 是 CampaignService 的具体实现
type campaignService struct {
	repo repositories.CampaignRepository
}

// NewCampaignService 创建一个新的 CampaignService 实例
func NewCampaignService(repo repositories.CampaignRepository) CampaignService {
	return &campaignService{repo: repo}
}

func (s *campaignService) CreateCampaign(advertiserID int64, name string, budget, bidAmount float64, bidType string) (*models.Campaign, error) {
	campaign := &models.Campaign{
		AdvertiserID: advertiserID,
		Name:         name,
		Budget:       budget,
		BidType:      bidType,
		BidAmount:    bidAmount,
		Status:       1, // 默认状态为投放中
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	err := s.repo.Create(campaign)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s *campaignService) ListCampaigns() ([]models.Campaign, error) {
	return s.repo.FindAll()
}

func (s *campaignService) GetCampaignByID(id int64) (*models.Campaign, error) {
	campaign, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s *campaignService) UpdateCampaign(campaign *models.Campaign) (*models.Campaign, error) {
	// 检查活动是否存在
	existing, err := s.repo.FindByID(campaign.ID)
	if err != nil {
		return nil, errors.New("campaign not found")
	}

	// 更新必要的字段
	existing.Name = campaign.Name
	existing.Budget = campaign.Budget
	existing.BidType = campaign.BidType
	existing.BidAmount = campaign.BidAmount
	existing.Status = campaign.Status
	existing.UpdateTime = time.Now()

	err = s.repo.Update(existing)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *campaignService) DeleteCampaign(id int64) error {
	// 检查活动是否存在
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("campaign not found")
	}
	return s.repo.Delete(id)
}
