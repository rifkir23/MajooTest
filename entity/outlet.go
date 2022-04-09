package entity

import "time"

type Outlet struct {
	Id         int `gorm:"primary_key:auto_increment"`
	MerchantId int
	OutletName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CreatedBy  int
	UpdatedBy  int

	Merchant Merchant `gorm:"foreignkey:Id;references:MerchantId;" json:"merchant"`
}
