package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ntquang/ecommerce/internal/services"
	"github.com/ntquang/ecommerce/internal/vo"
	"github.com/ntquang/ecommerce/response"
)

// type UserController struct {
// 	UserService *services.UserService
// }

// func NewUserController() *UserController {
// 	return &UserController{
// 		UserService: services.NewUserService(),
// 	}
// }

// func (uc *UserController) GetUser(c *gin.Context) {
// 	// name := c.Param("name")
// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	"message": "user " + uc.UserService.GetInfoUser(),
// 	// 	"name":    "is " + name,
// 	// })
// 	// c.JSON(http.StatusOK, response.Response{
// 	// 	Code:    20003,
// 	// 	Message: "Sucess",
// 	// 	Data:    uc.UserService.GetInfoUser(),
// 	// })
// 	// response.SuccessResponse(c, 20001, services.NewUserService().GetInfoUser())
// 	response.SuccessResponse(c, 20001, uc.UserService.GetInfoUser())
// }

type UserController struct {
	userService services.IUserservice
}

func NewUserController(
	userService services.IUserservice,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistrationRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, "", fmt.Errorf(err.Error()))
	}
	fmt.Printf("Email params::\n%v", params.Email)
	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil, "")
}
