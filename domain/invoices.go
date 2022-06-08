package domain

import (
	"gorm.io/gorm"

	"time"
)

type PaymentType string

const (
	GOPAY     = "GOPAY"
	OVO       = "OVO"
	SHOPEEPAY = "SHOPEEPAY"
)

type PayType string

const (
	ROOMS = "ROOMS"
	FOODS = "FOODS"
)

type Invoices struct {
	ID          uint           `gorm:"primaryKey;autoIncrementIncrement"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	ResidentID  uint
	RoomID      uint
	TotalPrice  uint
	PayDate     time.Time
	PaymentType PaymentType
	PayType     PayType
}
