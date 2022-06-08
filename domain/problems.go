package domain

import (
	"gorm.io/gorm"

	"time"
)

type ProblemType string

const (
	TIDAKPENTINGTIDAKGAWAT = "TIDAKPENTINGTIDAKGAWAT"
	PENTINGTIDAKGAWAT      = "PENTINGTIDAKGAWAT"
	PENTINGGAWAT           = "PENTINGGAWAT"
)

type Problems struct {
	ID          uint           `gorm:"primaryKey;autoIncrementIncrement"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	RoomID      uint
	ProblemType ProblemType
	Description string
}
