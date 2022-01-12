package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wilopo-cargo/microservice-receipt-sea/dto"
	"github.com/wilopo-cargo/microservice-receipt-sea/helper"
	"github.com/wilopo-cargo/microservice-receipt-sea/service"
	"net/http"
)

//ReceiptSeaController is a ...
type ReceiptSeaController interface {
	All(context *gin.Context)
	FindByNumber(context *gin.Context)
	Count(context *gin.Context)
	Delay(context *gin.Context)
	ArrivedSoon(context *gin.Context)
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

// All godoc
// @Summary All example
// @Schemes
// @Description All Receipt Sea
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response{data=[]entity.Resi}
// @Router /all [GET]
func (c *receiptSeaController) All(context *gin.Context) {
	var receipts = c.receiptSeaService.All()
	res := helper.BuildResponse(true, "OK", receipts)
	context.JSON(http.StatusOK, res)
}

// FindByNumber godoc
// @Summary All example
// @Schemes
// @Description Receipt Find By Number
// @Accept json
// @Produce json
// @Param FindByNumber body dto.ReceiptSeaJson true "FindByNumber"
// @Success 200 {object} helper.Response{data=[]entity.Resi}
// @Router /find [POST]

func (c *receiptSeaController) FindByNumber(context *gin.Context) {
	var receiptSeaJson dto.ReceiptSeaJson
	context.BindJSON(&receiptSeaJson)

	var receipt_sea = c.receiptSeaService.FindByNumber(receiptSeaJson.ReceiptSeaNumber)
	//if (receipt_sea == entity.Resi{}) {
	//	res := helper.BuildErrorResponse("Data Not Found", "No data with given id", helper.EmptyObj{})
	//	context.JSON(http.StatusNotFound, res)
	//} else {
	res := helper.BuildResponse(true, "OK", receipt_sea)
	context.JSON(http.StatusOK, res)
	//}
}

// Count godoc
// @Summary All example
// @Schemes
// @Description Count Receipt Sea
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response{data=dto.CountDTO}
// @Router /count [GET]
func (c *receiptSeaController) Count(context *gin.Context) {
	var countReceiptDTO dto.CountDTO
	result := c.receiptSeaService.Count(countReceiptDTO)
	res := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, res)
}

func (c *receiptSeaController) Delay(context *gin.Context) {
	var delayDto dto.DelayDTO
	var receipts = c.receiptSeaService.Delay(delayDto)
	res := helper.BuildResponse(true, "OK", receipts)
	context.JSON(http.StatusOK, res)
}

func (c *receiptSeaController) ArrivedSoon(context *gin.Context) {
	var arrivedSoonDto dto.ArrivedSoonDTO
	var receipts = c.receiptSeaService.ArrivedSoon(arrivedSoonDto)
	res := helper.BuildResponse(true, "OK", receipts)
	context.JSON(http.StatusOK, res)
}
