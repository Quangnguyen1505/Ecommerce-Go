package account

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ntquang/ecommerce/internal/model"
	"github.com/ntquang/ecommerce/internal/services"
	"github.com/ntquang/ecommerce/internal/utlis/context"
	"github.com/ntquang/ecommerce/response"
)

var User2fa = new(cUser2FA)

type cUser2FA struct{}

// User Login 2FA Documentation
// @Summary      User 2FA
// @Description  When user is login after 2 factor authentication
// @Tags         account 2fa
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authorization token"
// @Param        payload body model.SetupTwoFactorAuthInput true "payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.ErrResponse
// @Router       /user/two-factor/setup [post]
func (c *cUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params model.SetupTwoFactorAuthInput
	if err := ctx.ShouldBind(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrTwoFactorAuthSetUpFailed, "Missing or Invalid SetUp Two factor Authentication", err)
		return
	}
	//get userId from uuid (token)
	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrTwoFactorAuthSetUpFailed, "Missing get UUID", err)
	}
	log.Println("UserId::", userId)
	params.UserId = uint32(userId)
	resultCode, err := services.UserLogin().SetupTwoFactorAuth(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrTwoFactorAuthSetUpFailed, "Setup Two Factor Auth failed", err)
		return
	}
	response.SuccessResponse(ctx, resultCode, nil, "")
}
