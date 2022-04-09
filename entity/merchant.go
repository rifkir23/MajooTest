package entity

import "time"

type Merchant struct {
	Id           int `gorm:"primary_key:auto_increment"`
	UserId       int
	MerchantName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CreatedBy    int
	UpdatedBy    int

	User User `gorm:"foreignkey:Id;references:UserId;" json:"user"`
}
