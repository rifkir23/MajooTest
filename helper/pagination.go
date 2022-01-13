package helper

import (
	"github.com/wilopo-cargo/microservice-receipt-sea/dto"

	"gorm.io/gorm"
)

type Pagination interface {
	Paginate(b dto.BodyListReceipt) *gorm.DB
}

func Paginate(b dto.BodyListReceipt) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		page := b.Page
		if page == 0 {
			page = 1
		}

		pageSize := b.Limit
		switch {
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
