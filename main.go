package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rifkir23/MjTest/config"
	"github.com/rifkir23/MjTest/controller"
	docs "github.com/rifkir23/MjTest/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

var (
	db          = config.SetupDatabaseConnection()
	transaction = controller.NewTransactionController()
	auth        = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	r.GET("/login", auth.Login)

	invoice := r.Group("/transaction")
	{
		invoice.GET("report-outlet", transaction.TransactionReportByOutlet)
		invoice.GET("report-merchant", transaction.TransactionReportByMerchant)
	}

	docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_BASE_PATH")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
