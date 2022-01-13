package helper

import (
	"gorm.io/gorm"
)

type Pagination interface {
	Paginate(page int64, limit int64) *gorm.DB
}

func PaginateReceipt(page int64, limit int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		page := page
		if page == 0 {
			page = 1
		}

		pageSize := limit
		switch {
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
