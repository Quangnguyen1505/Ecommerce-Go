package account

import (
	"github.com/gin-gonic/gin"
	"github.com/ntquang/ecommerce/global"
	"github.com/ntquang/ecommerce/internal/model"
	"github.com/ntquang/ecommerce/internal/services"
	"github.com/ntquang/ecommerce/response"
	"go.uber.org/zap"
)

var Login = new(cUserLogin)

type cUserLogin struct{}

// User Login Documentation
// @Summary      User Login
// @Description  When user is login save dbs
// @Tags         account manager
// @Accept       json
// @Produce      json
// @Param        payload body model.RegisterInput true "payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.ErrResponse
// @Router       /user/login [post]
func (c *cUserLogin) Login(ctx *gin.Context) {
	//implement login for login
	err := services.UserLogin().Login(ctx)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}

// User Registrasion Documentation
// @Summary      User Registrasion
// @Description  When user is registered send otp to email
// @Tags         account manager
// @Accept       json
// @Produce      json
// @Param        payload body model.RegisterInput true "payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.ErrResponse
// @Router       /user/register [post]
func (c *cUserLogin) Register(ctx *gin.Context) {
	//implement login for register
	var params model.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
	}
	statusCode, err := services.UserLogin().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Error registering user OTP", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, statusCode, nil)
}

// User Verify OTP Documentation
// @Summary      User Verify OTP
// @Description  When user is register after verify otp
// @Tags         account manager
// @Accept       json
// @Produce      json
// @Param        payload body model.VerifyInput true "payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.ErrResponse
// @Router       /user/verifyOTP [post]
func (c *cUserLogin) VerifyOTP(ctx *gin.Context) {
	var params model.VerifyInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
	}
	result, err := services.UserLogin().VerifyOtp(ctx, &params)
	if err != nil {
		global.Logger.Error("Error verify user OTP", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}
