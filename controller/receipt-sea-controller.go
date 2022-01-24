package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wilopo-cargo/microservice-receipt-sea/dto"
	"github.com/wilopo-cargo/microservice-receipt-sea/helper"
	"github.com/wilopo-cargo/microservice-receipt-sea/service"
	"github.com/wilopo-cargo/microservice-receipt-sea/utility"
	"net/http"
	"strconv"
)

//ReceiptSeaController is a ...
type ReceiptSeaController interface {
	Count(context *gin.Context)
	Detail(context *gin.Context)
	List(context *gin.Context)
	ReceiptByContainer(context *gin.Context)
}

type receiptSeaController struct {
	receiptSeaService service.ReceiptSeaService
	jwtService        service.JWTService
}

//NewReceiptSeaController create a new instances of BoookController
func NewReceiptSeaController(receiptSeaServ service.ReceiptSeaService, jwtServ service.JWTService) ReceiptSeaController {
	return &receiptSeaController{
		receiptSeaService: receiptSeaServ,
		//jwtService:  jwtServ,
	}
}

// Count godoc
// @Summary All example
// @Schemes
// @Description Count Receipt Sea
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token. Format: Bearer access_token"
// @Success 200 {object} helper.Response{data=dto.CountReceiptSea}
// @Router /count [GET]
func (c *receiptSeaController) Count(context *gin.Context) {
	tokenAuth, errToken := utility.ValidateJwtToken(context.Request)
	if errToken != nil {
		panic(errToken.Error())
	}

	var countReceiptDTO dto.CountReceiptSea
	result := c.receiptSeaService.Count(int64(tokenAuth.UserId), countReceiptDTO)
	res := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, res)
}

// Detail godoc
// @Summary All example
// @Schemes
// @Description Receipt Detail
// @Accept json
// @Produce json
// @Param Detail receiptId  query int  false  "ReceiptId"
// @Param Detail containerId  query int  false  "ContainerId"
// @Success 200 {object} helper.Response{data=dto.ReceiptListResult}
// @Router /detail [GET]
func (c *receiptSeaController) Detail(context *gin.Context) {
	receiptId, err := strconv.ParseInt(context.Query("receiptId"), 0, 0)
	containerId, err := strconv.ParseInt(context.Query("containerId"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param int was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var receiptSea = c.receiptSeaService.Detail(receiptId, containerId)

	res := helper.BuildResponse(true, "OK", receiptSea)
	context.JSON(http.StatusOK, res)
}

// List godoc
// @Summary All example
// @Schemes
// @Description Receipt List
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token. Format: Bearer access_token"
// @Param List page  query int  false  "Pages"
// @Param List limit  query int  false  "Limit"
// @Param List status query  string  false  "Status"
// @Success 200 {object} helper.Response{data=dto.ReceiptListResult}
// @Router /list [GET]
func (c *receiptSeaController) List(context *gin.Context) {
	tokenAuth, errToken := utility.ValidateJwtToken(context.Request)
	if errToken != nil {
		panic(errToken.Error())
	}

	page, err := strconv.ParseInt(context.Query("page"), 0, 0)
	limit, err := strconv.ParseInt(context.Query("limit"), 0, 0)
	status := context.Query("status")
	if err != nil {
		res := helper.BuildErrorResponse("No param int was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var receipts = c.receiptSeaService.List(int64(tokenAuth.UserId), page, limit, status)
	res := helper.BuildResponse(true, "OK", receipts)
	context.JSON(http.StatusOK, res)
}

// ReceiptByContainer godoc
// @Summary All example
// @Schemes
// @Description Receipt By Container
// @Accept json
// @Produce json
// @Param ReceiptByContainer body dto.ReceiptNumber true "ReceiptByContainer"
// @Success 200 {object} helper.Response{data=dto.ContainerByReceipt}
// @Router /container-by-receipt [POST]
func (c *receiptSeaController) ReceiptByContainer(context *gin.Context) {
	var receiptNumber dto.ReceiptNumber
	context.BindJSON(&receiptNumber)

	var receiptSea = c.receiptSeaService.ReceiptByContainer(receiptNumber.ReceiptSeaNumber)
	res := helper.BuildResponse(true, "OK", receiptSea)
	context.JSON(http.StatusOK, res)
}
