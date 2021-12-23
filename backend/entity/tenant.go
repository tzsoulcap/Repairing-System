package entity

import (
	"time"

	"gorm.io/gorm"
)

type Employer struct {
	gorm.Model

	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Tenants  []Tenant `gorm:"foreignKey:OwnerID"`
	//Promotion
	PromotionActives []PromotionActive `gorm:"foreignKey:EmployerID"`
}

type Tenant struct {
	gorm.Model
	Name     string
	Idcard   string `gorm:"uniqueIndex"`
	Tel      string
	BirthDay time.Time
	Age      int
	Address  uint
	Village  string
	Email    string `gorm:"uniqueIndex"`
	Password string
	OwnerID  *uint
	Owner    Employer

	PrefixID *uint
	Prefix   Prefix `gorm:"references:id"`

	CareerID *uint
	Career   Career `gorm:"references:id"`

	GenderID *uint
	Gender   Gender `gorm:"references:id"`

	CantonID *uint
	Canton   Canton `gorm:"references:id"`
	//Rental
	RENTALS []Rental `gorm:"foreignKey:TenantID"`
	//Repair
	Repairs []Repair `gorm:"foreignKey:TenantID"`
	//ReturnRoom
	ReturnRoom []ReturnRoom `gorm:"foreignKey:TenantID"`
}

type Prefix struct {
	gorm.Model

	Name    string
	Tenants []Tenant `gorm:"foreignKey:PrefixID"`
}

type Career struct {
	gorm.Model

	Name    string
	Tenants []Tenant `gorm:"foreignKey:CareerID"`
}

type Gender struct {
	gorm.Model

	Name    string
	Tenants []Tenant `gorm:"foreignKey:GenderID"`
}

type Province struct {
	gorm.Model
	Name        string
	Prefectures []Prefecture `gorm:"foreignKey:ProvinceID"`
}

type Prefecture struct {
	gorm.Model
	Name       string
	Cantons    []Canton `gorm:"foreignKey:PrefectureID"`
	ProvinceID *uint
	Province   Province `gorm:"references:id"`
}

type Canton struct {
	gorm.Model
	Name         string
	Tenants      []Tenant `gorm:"foreignKey:CantonID"`
	PrefectureID *uint
	Prefecture   Prefecture `gorm:"references:id"`
}
