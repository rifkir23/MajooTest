package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wilopo-cargo/microservice-receipt-sea/config"
	"github.com/wilopo-cargo/microservice-receipt-sea/controller"
	"github.com/wilopo-cargo/microservice-receipt-sea/repository"
	"github.com/wilopo-cargo/microservice-receipt-sea/service"
	"gorm.io/gorm"
)

var (
	db                   *gorm.DB                        = config.SetupDatabaseConnection()
	receiptSeaRepository repository.ReceiptSeaRepository = repository.NewReceiptSeaRepository(db)
	jwtService           service.JWTService              = service.NewJWTService()
	receiptSeaService    service.ReceiptSeaService       = service.NewReceiptSeaService(receiptSeaRepository)
	receiptSeaController controller.ReceiptSeaController = controller.NewReceiptSeaController(receiptSeaService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	r.POST("/main", receiptSeaController.FindByNumber)
	r.GET("/all", receiptSeaController.All)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
