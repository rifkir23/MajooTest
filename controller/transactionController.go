package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rifkir23/MjTest/helper"
	"github.com/rifkir23/MjTest/service"
	"github.com/sirupsen/logrus"
	"strconv"
)

var (
	transactionService = service.NewTransactionService()
)

type TransactionController interface {
	TransactionReportByOutlet(ctx *gin.Context)
	TransactionReportByMerchant(ctx *gin.Context)
}

func NewTransactionController() TransactionController {
	return &transactionController{}
}

type transactionController struct {
}

// TransactionReportByOutlet godoc
// @Summary Transaction Report By Outlet
// @Schemes
// @Description Transaction Report By Outlet
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token"
// @Param List page  query int  false  "page"
// @Param List limit  query int  false  "limit"
// @Security BearerToken
// @Success 200 {object} helper.PaginationResponse{}
// @Router /transaction/report-outlet [GET]
func (t transactionController) TransactionReportByOutlet(ctx *gin.Context) {
	/*Validate Token*/
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := jwtService.ValidateToken(authHeader)
	if errToken != nil {
		logrus.Error(errToken)
	}
	claims := token.Claims.(jwt.MapClaims)
	userIdStr := fmt.Sprintf("%v", claims["user_id"])
	userId, errAt := strconv.Atoi(userIdStr)
	if errAt != nil {
		logrus.Error(errToken)
	}

	/*Pagination*/
	pagination := helper.GeneratePaginationFromRequest(ctx)

	var transactions = transactionService.TransactionReportByOutlet(pagination, userId)
	helper.ResponseSuccess(transactions, ctx)
}

// TransactionReportByMerchant godoc
// @Summary Transaction Report By Merchant
// @Schemes
// @Description Transaction Report By Merchant
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token"
// @Param List page  query int  false  "page"
// @Param List limit  query int  false  "limit"
// @Security BearerToken
// @Success 200 {object} helper.PaginationResponse{}
// @Router /transaction/report-merchant [GET]
func (t transactionController) TransactionReportByMerchant(ctx *gin.Context) {
	/*Validate Token*/
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := jwtService.ValidateToken(authHeader)
	if errToken != nil {
		logrus.Error(errToken)
	}
	claims := token.Claims.(jwt.MapClaims)
	userIdStr := fmt.Sprintf("%v", claims["user_id"])
	userId, errAt := strconv.Atoi(userIdStr)
	if errAt != nil {
		logrus.Error(errToken)
	}

	/*Pagination*/
	pagination := helper.GeneratePaginationFromRequest(ctx)

	var transactions = transactionService.TransactionReportByMerchant(pagination, userId)
	helper.ResponseSuccess(transactions, ctx)
}
