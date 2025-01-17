package services

import (
	"errors"
	"time"

	"mini-dsp-backend/models"
	"mini-dsp-backend/repositories"
)

// CreativeService 定义了创意管理的业务接口
type CreativeService interface {
	CreateCreative(campaignID int64, creativeName, creativeType, landingPageUrl string) (*models.Creative, error)
	ListCreatives() ([]models.Creative, error)
	GetCreativeByID(id int64) (*models.Creative, error)
	UpdateCreative(creative *models.Creative) (*models.Creative, error)
	DeleteCreative(id int64) error
}

// creativeService 是 CreativeService 的具体实现
type creativeService struct {
	repo repositories.CreativeRepository
}

// NewCreativeService 创建一个新的 CreativeService 实例
func NewCreativeService(repo repositories.CreativeRepository) CreativeService {
	return &creativeService{repo: repo}
}

func (s *creativeService) CreateCreative(campaignID int64, creativeName, creativeType, landingPageUrl string) (*models.Creative, error) {
	creative := &models.Creative{
		CampaignID:     campaignID,
		CreativeName:   creativeName,
		CreativeType:   creativeType,
		LandingPageUrl: landingPageUrl,
		Status:         1, // 默认启用
		CreateTime:     time.Now(),
		UpdateTime:     time.Now(),
	}
	err := s.repo.Create(creative)
	if err != nil {
		return nil, err
	}
	return creative, nil
}

func (s *creativeService) ListCreatives() ([]models.Creative, error) {
	return s.repo.FindAll()
}

func (s *creativeService) GetCreativeByID(id int64) (*models.Creative, error) {
	creative, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return creative, nil
}

func (s *creativeService) UpdateCreative(creative *models.Creative) (*models.Creative, error) {
	existing, err := s.repo.FindByID(creative.ID)
	if err != nil {
		return nil, errors.New("creative not found")
	}

	existing.CreativeName = creative.CreativeName
	existing.CreativeType = creative.CreativeType
	existing.LandingPageUrl = creative.LandingPageUrl
	existing.Status = creative.Status
	existing.UpdateTime = time.Now()

	err = s.repo.Update(existing)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *creativeService) DeleteCreative(id int64) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("creative not found")
	}
	return s.repo.Delete(id)
}
