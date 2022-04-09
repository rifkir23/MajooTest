package helper

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type Errors struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type PaginationResponse struct {
	Pagination Pagination  `json:"pagination"`
	Content    interface{} `json:"data"`
}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}

func ResponseSuccess(data interface{}, ctx *gin.Context) {
	res := BuildResponse(true, "success", data)
	ctx.JSON(http.StatusOK, res)
}

func DialogSuccess(message string, ctx *gin.Context) {
	res := BuildResponse(true, message, nil)
	ctx.JSON(http.StatusOK, res)
}

func DialogError(message string, httpStatus int, ctx *gin.Context) {
	res := BuildResponse(true, message, nil)
	//ctx.JSON(httpStatus, res)
	ctx.AbortWithStatusJSON(httpStatus, res)
}

func BuildPaginationResponse(dto interface{}, pagination Pagination) PaginationResponse {
	return PaginationResponse{
		Pagination: pagination,
		Content:    dto,
	}
}
