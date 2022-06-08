package domain

import (
	"gorm.io/gorm"

	"time"
)

type RoomStatus string

const (
	UNOCCUPIED = "UNOCCUPIED"
	OCCUPIED   = "OCCUPIED"
	NOTREADY   = "NOTREADY"
)

type Rooms struct {
	ID                uint           `gorm:"primaryKey;autoIncrementIncrement"`
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	ResidentID        uint
	Price             uint
	StartOccupiedDate time.Time
	EndOccupiedDate   time.Time
	Status            RoomStatus
	QRAPICode         string
}
