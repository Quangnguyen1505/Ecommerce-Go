package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"messgae"`
	Data    interface{} `json:"data"`
}

type ErrResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"messgae"`
	Detail  interface{} `json:"detail"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}, message string) {
	if message == "" {
		message = msg[code]
	}
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string, err error) {
	if message == "" {
		message = msg[code]
	}
	c.JSON(http.StatusOK, ErrResponse{
		Code:    code,
		Message: message,
		Detail:  err,
	})
}
