package models

import "time"

const (
	AdvertiserClose = 0
	AdvertiserOpen  = 1
)

type Advertiser struct {
	ID         int64  `gorm:"primaryKey"`
	Name       string `gorm:"size:255"`
	Contact    string `gorm:"size:255"`
	Status     int    // 0:禁用, 1:可用
	CreateTime time.Time
	UpdateTime time.Time
}
