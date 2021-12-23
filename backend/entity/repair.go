package entity

import (
	"time"

	"gorm.io/gorm"
)

type TypeEquipment struct {
	gorm.Model
	Name       string
	Equipments []Equipment `gorm:"foreignKey:TypeEquipmentID"`
}

type Equipment struct {
	gorm.Model
	Name            string
	TypeEquipmentID *uint
	TypeEquipment   TypeEquipment   `gorm:"references:id"`
	RoomEquipments  []RoomEquipment `gorm:"foreignKey:EquipmentID"`
}

type RoomEquipment struct {
	gorm.Model
	RoomID      *uint
	Room        Room `gorm:"references:id"`
	EquipmentID *uint
	Equipment   Equipment `gorm:"references:id"`
	Repairs     []Repair  `gorm:"foreignKey:RoomEquipmentID"`
}

type Repair struct {
	gorm.Model
	AddedTime       time.Time
	Note            string
	TenantID        *uint
	Tenant          Tenant `gorm:"references:id"`
	RoomEquipmentID *uint
	RoomEquipment   RoomEquipment `gorm:"references:id"`
	RentalID        *uint
	Rental          Rental `gorm:"references:id"`
}
