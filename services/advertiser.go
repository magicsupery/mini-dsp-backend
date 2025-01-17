package services

import (
	"mini-dsp-backend/models"
	"mini-dsp-backend/repositories"
	"time"
)

// AdvertiserService 定义了广告主管理的服务接口
type AdvertiserService interface {
	CreateAdvertiser(name, contact string) (*models.Advertiser, error)
	ListAdvertisers() ([]models.Advertiser, error)
	GetAdvertiserByID(id int64) (*models.Advertiser, error)
	UpdateAdvertiser(ad *models.Advertiser) (*models.Advertiser, error)
	DeleteAdvertiser(id int64) error
}

// advertiserService 是 AdvertiserService 接口的具体实现
type advertiserService struct {
	repo repositories.AdvertiserRepository
}

// NewAdvertiserService 创建一个新的 AdvertiserService 实例
func NewAdvertiserService(repo repositories.AdvertiserRepository) AdvertiserService {
	return &advertiserService{repo: repo}
}

// CreateAdvertiser 创建一个新的广告主
func (s *advertiserService) CreateAdvertiser(name, contact string) (*models.Advertiser, error) {
	ad := &models.Advertiser{
		Name:       name,
		Contact:    contact,
		Status:     1,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := s.repo.Create(ad)
	return ad, err
}

// ListAdvertisers 列出所有广告主
func (s *advertiserService) ListAdvertisers() ([]models.Advertiser, error) {
	return s.repo.FindAll()
}

// GetAdvertiserByID 通过 ID 获取广告主
func (s *advertiserService) GetAdvertiserByID(id int64) (*models.Advertiser, error) {
	return s.repo.FindByID(id)
}

// UpdateAdvertiser 更新广告主信息
func (s *advertiserService) UpdateAdvertiser(ad *models.Advertiser) (*models.Advertiser, error) {
	err := s.repo.Update(ad)
	return ad, err
}

// DeleteAdvertiser 删除指定 ID 的广告主
func (s *advertiserService) DeleteAdvertiser(id int64) error {
	return s.repo.Delete(id)
}
