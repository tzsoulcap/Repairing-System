package entity

import (
	"time"

	"gorm.io/gorm"
)

type PromotionType struct {
	gorm.Model
	Name           string
	PromotionInfos []PromotionInfo `gorm:"foreignKey:TypeID"`
}

type Festival struct {
	gorm.Model
	Name             string
	PromotionActives []PromotionActive `gorm:"foreignKey:FestID"`
}

type PromotionInfo struct {
	gorm.Model
	Name             string
	TypeID           *uint
	Type             PromotionType
	PromotionActives []PromotionActive `gorm:"foreignKey:InfoID"`
}

type PromotionActive struct {
	gorm.Model
	DurationStart time.Time
	DurationEnd   time.Time

	EmployerID *uint
	Employer   Employer

	FestID *uint
	Fest   Festival

	InfoID *uint
	Info   PromotionInfo
}
