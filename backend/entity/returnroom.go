package entity

import (
	"time"

	"gorm.io/gorm"
)

type ReturnRoom struct {
	gorm.Model
	MoveOutTime time.Time

	TenantID *uint
	Tenant   Tenant `gorm:"references:id"`

	RoomID *uint
	Room   Room `gorm:"references:id"`

	RentalID *uint
	Rental   Rental `gorm:"references:id"`
}
