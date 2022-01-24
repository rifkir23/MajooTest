package repository

import (
	"github.com/wilopo-cargo/microservice-receipt-sea/dto"
	"github.com/wilopo-cargo/microservice-receipt-sea/entity"
	"github.com/wilopo-cargo/microservice-receipt-sea/helper"
	"gorm.io/gorm"
)

type ReceiptRepository interface {
	List(customerId int64, page int64, limit int64, receiptType string) dto.ReceiptListByTypeResult
}

type receiptConnection struct {
	connection *gorm.DB
}

//NewReceiptRepository creates an instance ReceiptRepository
func NewReceiptRepository(dbConn *gorm.DB) ReceiptRepository {
	return &receiptConnection{
		connection: dbConn,
	}
}

func (db *receiptConnection) List(customerId int64, page int64, limit int64, receiptType string) dto.ReceiptListByTypeResult {
	var receiptSea entity.Resi
	var receiptAir entity.ReceiptAir
	var receiptList []dto.ReceiptListByType
	var pagination dto.Pagination
	var countList int64

	if receiptType == "sea" {
		db.connection.Model(&receiptSea).Select("id_resi as ReceiptId,tanggal as Date,'"+receiptType+"'as Type,nomor as ReceiptNumber,IF(konfirmasi_resi = 0,0,1 ) as Status").
			Where("cust_id", customerId).
			Count(&countList).Scopes(helper.PaginateReceipt(page, limit)).Find(&receiptList)
	} else if receiptType == "air" {
		db.connection.Model(&receiptAir).Select("id_resi_udara as ReceiptId,tanggal_resi as Date,'"+receiptType+"'as Type,nomor_resi as ReceiptNumber,IF(id_invoice > 0,1,0 ) as Status").
			Where("id_cust", customerId).
			Count(&countList).Scopes(helper.PaginateReceipt(page, limit)).Find(&receiptList)
	}

	/*Current Page*/
	if page >= (countList / limit) {
		pagination.CurrentPage = countList / limit
	} else {
		pagination.CurrentPage = page
	}
	/*Prev Page*/
	if (page-1) > 0 && (countList/limit) > 0 {
		if page >= (countList / limit) {
			pagination.PrevPage = pagination.CurrentPage - 1
		} else {
			pagination.PrevPage = page - 1
		}
	}
	/*Next Page*/
	if (page + 1) < (countList / limit) {
		pagination.NextPage = page + 1
	} else {
		pagination.NextPage = 0
	}
	pagination.TotalElement = countList
	pagination.TotalPage = countList / limit

	results := dto.ReceiptListByTypeResult{
		Pagination: pagination,
		Receipt:    receiptList,
	}

	return results
}
