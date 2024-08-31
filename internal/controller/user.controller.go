package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ntquang/ecommerce/internal/services"
	"github.com/ntquang/ecommerce/response"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: services.NewUserService(),
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	// name := c.Param("name")
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "user " + uc.UserService.GetInfoUser(),
	// 	"name":    "is " + name,
	// })
	// c.JSON(http.StatusOK, response.Response{
	// 	Code:    20003,
	// 	Message: "Sucess",
	// 	Data:    uc.UserService.GetInfoUser(),
	// })
	// response.SuccessResponse(c, 20001, services.NewUserService().GetInfoUser())
	response.SuccessResponse(c, 20001, uc.UserService.GetInfoUser())
}
