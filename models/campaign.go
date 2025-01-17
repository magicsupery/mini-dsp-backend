package models

import "time"

type Campaign struct {
	ID           int64 `gorm:"primaryKey"`
	AdvertiserID int64
	Name         string `gorm:"size:255"`
	Budget       float64
	BidType      string `gorm:"size:50"` // CPC, CPM, oCPA...
	BidAmount    float64
	Status       int // 0:暂停, 1:投放中, 2:已下线
	CreateTime   time.Time
	UpdateTime   time.Time
}
