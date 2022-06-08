package domain

import (
	"gorm.io/gorm"

	"time"
)

type Admins struct {
	ID        uint           `gorm:"primaryKey;autoIncrementIncrement"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	AdminName string
	Username  string
	Password  string
}
