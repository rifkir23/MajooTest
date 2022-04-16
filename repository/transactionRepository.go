package repository

import (
	"github.com/rifkir23/MjTest/dto"
	"github.com/rifkir23/MjTest/entity"
	"github.com/rifkir23/MjTest/helper"
)

type TransactionRepository interface {
	TransactionReportByOutlet(pagination helper.Pagination, userId int) (*helper.Pagination, error)
	TransactionReportByMerchant(pagination helper.Pagination, userId int, totalRows int, startDate string, endDate string) (*helper.Pagination, error)
}

type transactionRepo struct {
}

func NewTransactionRepo() TransactionRepository {
	return &transactionRepo{}
}

func (t transactionRepo) TransactionReportByOutlet(pagination helper.Pagination, userId int) (*helper.Pagination, error) {
	var transactions []*entity.Transaction
	var transactionByOutlet []*dto.TransactionReportByOutlet
	var totalRows int64

	query := db.Model(transactions).
		Select("merchant_name,outlet_name,sum(transactions.bill_total) as BillTotal,transactions.created_at as TransactionDate").
		Joins("Outlet").
		Joins("LEFT JOIN merchants on Outlet.merchant_id = merchants.id").
		Where("(DATE(transactions.created_at) BETWEEN ? AND ?)", "2021-11-01", "2021-11-30").
		Where("user_id", userId).
		Group("outlet_id,Date(transactions.created_at)")

	query.Count(&totalRows)
	query.Scopes(helper.PaginateQuery(&transactionByOutlet, &pagination, totalRows)).Find(&transactionByOutlet)

	pagination.Content = transactionByOutlet
	return &pagination, nil
}

func (t transactionRepo) TransactionReportByMerchant(pagination helper.Pagination, userId int, totalRows int, startDate string, endDate string) (*helper.Pagination, error) {
	var transactions []*entity.Transaction
	var transactionByMerchant []*dto.TransactionReportByMerchant

	query := db.Model(transactions).
		Select("merchant_name,sum(transactions.bill_total) as BillTotal,transactions.created_at as TransactionDate").
		Joins("Outlet").
		Joins("LEFT JOIN merchants on Outlet.merchant_id = merchants.id").
		Where("(DATE(transactions.created_at) BETWEEN ? AND ?)", startDate, endDate).
		Where("user_id", userId).
		Group("transactions.merchant_id,Date(transactions.created_at)")

	query.Scopes(helper.PaginateQuery(&transactionByMerchant, &pagination, int64(totalRows))).Find(&transactionByMerchant)

	pagination.Content = transactionByMerchant
	return &pagination, nil
}
