package service

import (
	"github.com/jinzhu/copier"
	"github.com/rifkir23/MjTest/dto"
	"github.com/rifkir23/MjTest/helper"
	"github.com/rifkir23/MjTest/repository"
	"github.com/sirupsen/logrus"
)

var (
	repoTransaction = repository.NewTransactionRepo()
)

type TransactionService interface {
	TransactionReportByOutlet(pagination helper.Pagination, userId int) helper.PaginationResponse
	TransactionReportByMerchant(pagination helper.Pagination, userId int) helper.PaginationResponse
}

func NewTransactionService() TransactionService {
	return &transactionService{}
}

type transactionService struct {
}

func (t transactionService) TransactionReportByOutlet(pagination helper.Pagination, userId int) helper.PaginationResponse {
	transactions, err := repoTransaction.TransactionReportByOutlet(pagination, userId)
	if err != nil {
		logrus.Error(err)
	}

	responses := make([]dto.TransactionReportByOutlet, 0)
	data := transactions.Content.([]*dto.TransactionReportByOutlet)
	if len(data) != 0 {
		for _, transaction := range data {
			response := new(dto.TransactionReportByOutlet)
			err := copier.Copy(response, transaction)
			if err != nil {
				logrus.Error(err)
			}

			responses = append(responses, *response)
		}
	}
	return helper.BuildPaginationResponse(responses, *transactions)
}

func (t transactionService) TransactionReportByMerchant(pagination helper.Pagination, userId int) helper.PaginationResponse {
	transactions, err := repoTransaction.TransactionReportByMerchant(pagination, userId)
	if err != nil {
		logrus.Error(err)
	}

	responses := make([]dto.TransactionReportByMerchant, 0)
	data := transactions.Content.([]*dto.TransactionReportByMerchant)
	if len(data) != 0 {
		for _, transaction := range data {
			response := new(dto.TransactionReportByMerchant)
			err := copier.Copy(response, transaction)
			if err != nil {
				logrus.Error(err)
			}

			responses = append(responses, *response)
		}
	}
	return helper.BuildPaginationResponse(responses, *transactions)
}
