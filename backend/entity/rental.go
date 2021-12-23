package entity

import (
	"time"

	"gorm.io/gorm"
)

type Roomtype struct {
	gorm.Model
	ROOM_TYPE  string
	ROOM_PRICE uint
	ROOMS      []Room `gorm:"foreignKey:RoomtypeID"`
}
type Roomstate struct {
	gorm.Model
	STATE_ROOM string
	ROOMS      []Room `gorm:"foreignKey:RoomstateID"`
}
type Room struct {
	gorm.Model
	ROOM_NUMBER string `gorm:"uniqueIndex"`
	RoomtypeID  *uint
	Roomtype    Roomtype `gorm:"references:id"`
	RoomstateID *uint
	Roomstate   Roomstate `gorm:"references:id"`
	//Repairs
	RoomEquipments []RoomEquipment `gorm:"foreignKey:RoomID"`
	RENTALS        []Rental        `gorm:"foreignKey:RoomID"`
	//ReturnRoom
	ReturnRoom []ReturnRoom `gorm:"foreignKey:RoomID"`
}
type Rentalstate struct {
	gorm.Model
	RENTAL_STATE string
	RENTALS      []Rental `gorm:"foreignKey:RentalstateID"`
}
type Rental struct {
	gorm.Model
	DAYTIME       time.Time
	CHACKIN       time.Time
	TenantID      *uint
	Tenant        Tenant `gorm:"references:id"`
	RoomID        *uint
	Room          Room `gorm:"references:id"`
	RentalstateID *uint
	Rentalstate   Rentalstate `gorm:"references:id"`
	//Repairs
	Repairs []Repair `gorm:"foreignKey:RentalID"`
	//ReturnRoom
	Type       string
	ReturnRoom []ReturnRoom `gorm:"foreignKey:RentalID"`
}
