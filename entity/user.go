package entity

import (
	"time"
)

type User struct {
	Id        int `gorm:"primary_key:auto_increment"`
	Name      string
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy int
	UpdatedBy int
}
