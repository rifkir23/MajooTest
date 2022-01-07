package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wilopo-cargo/microservice-receipt-sea/entity"
	"github.com/wilopo-cargo/microservice-receipt-sea/helper"
	"github.com/wilopo-cargo/microservice-receipt-sea/service"
	"net/http"
)

//ReceiptSeaController is a ...
type ReceiptSeaController interface {
	All(context *gin.Context)
	FindByNumber(context *gin.Context)
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
	var receipts []entity.Resi = c.receiptSeaService.All()
	res := helper.BuildResponse(true, "OK", receipts)
	context.JSON(http.StatusOK, res)
}

// FindByNumber godoc
// @Summary All example
// @Schemes
// @Description Receipt Find By Number
// @Accept json
// @Produce json
// @Param FindByNumber body ReceiptSeaJson true "FindByNumber"
// @Success 200 {object} helper.Response{data=[]entity.Resi}
// @Router /main [POST]
func (c *receiptSeaController) FindByNumber(context *gin.Context) {
	var receiptSeaJson ReceiptSeaJson
	context.BindJSON(&receiptSeaJson)

	var receipt_sea entity.Resi = c.receiptSeaService.FindByNumber(receiptSeaJson.ReceiptSeaNumber)
	//if (receipt_sea == entity.Resi{}) {
	//	res := helper.BuildErrorResponse("Data Not Found", "No data with given id", helper.EmptyObj{})
	//	context.JSON(http.StatusNotFound, res)
	//} else {
	res := helper.BuildResponse(true, "OK", receipt_sea)
	context.JSON(http.StatusOK, res)
	//}
}

type ReceiptSeaJson struct {
	ReceiptSeaNumber string `json:"receipt_sea_number"`
}
