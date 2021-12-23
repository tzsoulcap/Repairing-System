package entity

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Paydate time.Time

	EmployerID *uint
	Employer   Employer

	TenantID *uint
	Tenant   Tenant

	RentalID *uint
	Rental   Rental

	PromotionActiveID *uint
	PromotionActive   PromotionActive
}
