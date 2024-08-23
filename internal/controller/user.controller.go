package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ntquang/ecommerce/internal/services"
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
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "user " + uc.UserService.GetInfoUser(),
		"name":    "is " + name,
	})
}
