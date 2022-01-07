package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wilopo-cargo/microservice-receipt-sea/config"
	"github.com/wilopo-cargo/microservice-receipt-sea/controller"
	docs "github.com/wilopo-cargo/microservice-receipt-sea/docs"
	"github.com/wilopo-cargo/microservice-receipt-sea/repository"
	"github.com/wilopo-cargo/microservice-receipt-sea/service"
	"gorm.io/gorm"
	"net/http"
	//docs "./docs"
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

	docs.SwaggerInfo.BasePath = "/"
	v1 := r.Group("/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
