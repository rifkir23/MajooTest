package repository

import (
	"github.com/rifkir23/MjTest/entity"
)

type MerchantRepository interface {
	GetById(merchantId int) *entity.Merchant
}

type merchantRepo struct {
}

func NewMerchantRepo() MerchantRepository {
	return &merchantRepo{}
}

func (m merchantRepo) GetById(merchantId int) *entity.Merchant {
	var merchants *entity.Merchant
	db.Model(&merchants).First(&merchants, merchantId)

	return merchants
}
