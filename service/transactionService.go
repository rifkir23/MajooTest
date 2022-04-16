package service

import (
	"github.com/jinzhu/copier"
	"github.com/rifkir23/MjTest/dto"
	"github.com/rifkir23/MjTest/helper"
	"github.com/rifkir23/MjTest/repository"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var (
	repoTransaction = repository.NewTransactionRepo()
	repoMerchant    = repository.NewMerchantRepo()
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
	var responses []*dto.TransactionReportByMerchant

	/*get merchant*/
	merchant := repoMerchant.GetById(userId)

	/*initialize 30 days data*/
	findStartDisplayDate := (pagination.CurrentPage * pagination.Limit) - (pagination.Limit - 1)
	findEndDisplayDate := pagination.CurrentPage * pagination.Limit
	startDisplayDate := "2021-11-" + strconv.Itoa(findStartDisplayDate)
	if findStartDisplayDate < 10 {
		startDisplayDate = "2021-11-0" + strconv.Itoa(findStartDisplayDate)
	}
	endDisplayDate := "2021-11-" + strconv.Itoa(findEndDisplayDate)
	if findEndDisplayDate < 10 {
		endDisplayDate = "2021-11-0" + strconv.Itoa(findEndDisplayDate)
	}

	format := "2006-01-02 15:04:05"
	calEndDate, _ := time.Parse(format, "2021-11-30 00:00:00")
	calStartDate, _ := time.Parse(format, "2021-11-01 00:00:00")
	diff := calEndDate.Sub(calStartDate)
	totalDays := int(diff.Hours() / 24)
	totalRows := totalDays + 1

	calEndDisplayDate, _ := time.Parse(format, endDisplayDate+" 00:00:00")
	calStartDisplayDate, _ := time.Parse(format, startDisplayDate+" 00:00:00")
	diffDisplay := calEndDisplayDate.Sub(calStartDisplayDate)
	totalDisplayDays := int(diffDisplay.Hours() / 24)

	for i := 0; i <= totalDisplayDays; i++ {
		dateTransaction := calStartDate.AddDate(0, 0, i)
		responses = append(responses, &dto.TransactionReportByMerchant{
			MerchantName:    merchant.MerchantName,
			BillTotal:       0,
			TransactionDate: dateTransaction,
		})
	}
	/*end initialize 30 days data*/

	transactions, err := repoTransaction.TransactionReportByMerchant(pagination, userId, totalRows, startDisplayDate, endDisplayDate)
	if err != nil {
		logrus.Error(err)
	}

	data := transactions.Content.([]*dto.TransactionReportByMerchant)
	if len(data) != 0 {
		for _, transaction := range data {
			for _, response := range responses {
				if transaction.TransactionDate.Format("2006-01-02") == response.TransactionDate.Format("2006-01-02") {
					response.BillTotal = transaction.BillTotal
				}
			}
		}
	}
	return helper.BuildPaginationResponse(responses, *transactions)
}
