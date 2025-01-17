package models

import "time"

type User struct {
	ID         int64  `gorm:"primaryKey"`
	Username   string `gorm:"size:255;unique"`
	Password   string `gorm:"size:255"` // 加密后的密码
	Role       string `gorm:"size:50"`
	CreateTime time.Time
	UpdateTime time.Time
}
