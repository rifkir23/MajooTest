package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rifkir23/MjTest/dto"
	"github.com/rifkir23/MjTest/entity"
	"github.com/rifkir23/MjTest/helper"
	"github.com/rifkir23/MjTest/service"
	"net/http"
	"strconv"
)

var (
	authService = service.NewAuthService()
	jwtService  = service.NewJWTService()
)

type AuthController interface {
	Login(ctx *gin.Context)
}

func NewAuthController() AuthController {
	return &authController{}
}

type authController struct {
}

/*Login Form Data*/
func (a *authController) Login(ctx *gin.Context) {
	var loginDTO dto.Login
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := authService.VerifyCredential(loginDTO.Username, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := jwtService.GenerateToken(strconv.Itoa(v.Id))
		response := helper.BuildResponse(true, "OK!", generatedToken)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
