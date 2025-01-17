package models

import "time"

type Creative struct {
	ID             int64 `gorm:"primaryKey"`
	CampaignID     int64
	CreativeName   string `gorm:"size:255"`
	CreativeType   string `gorm:"size:50"` // image / video
	LandingPageUrl string `gorm:"size:500"`
	Status         int    // 0:暂停, 1:启用
	CreateTime     time.Time
	UpdateTime     time.Time
}
