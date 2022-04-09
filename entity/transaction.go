package entity

import "time"

type Transaction struct {
	Id         int `gorm:"primary_key:auto_increment"`
	MerchantId int
	OutletId   int
	BillTotal  float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CreatedBy  int
	UpdatedBy  int

	Merchant Merchant `gorm:"foreignkey:Id;references:MerchantId;" json:"merchant"`
	Outlet   Outlet   `gorm:"foreignkey:Id;references:OutletId;" json:"outlet"`
}
