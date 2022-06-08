package domain

import (
	"gorm.io/gorm"

	"time"
)

type StayType string

const (
	TEMPORARY StayType = "TEMPORARY"
	PERMANENT StayType = "PERMANENT"
)

type Residents struct {
	ID          uint           `gorm:"primaryKey;autoIncrementIncrement"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	StayType    StayType
	Bills       uint
	PhoneNumber uint
	FullName    string
	Occupation  string
	Address     string
	NIKNumber   uint
	Username    string
	Password    string
}
