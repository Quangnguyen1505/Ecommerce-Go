package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ntquang/ecommerce/global"
	"github.com/ntquang/ecommerce/internal/repo"
	"github.com/ntquang/ecommerce/internal/utlis/crypto"
	"github.com/ntquang/ecommerce/internal/utlis/random"
	"github.com/ntquang/ecommerce/internal/utlis/sendto"
	"github.com/ntquang/ecommerce/response"
	"go.uber.org/zap"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepo.GetInfoUser()
// }

// INTERFACE VERSION
type IUserservice interface {
	Register(email string, purpose string) int
}

type userService struct { //dua ve private vi IUserservice se truc tiep lam viec voi controller
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(
	userRepo repo.IUserRepository,
	userAuthRepo repo.IUserAuthRepository,
) IUserservice {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {
	// 0. Hash email
	hashEmail := crypto.GetHash(email)
	fmt.Printf("Email after hash is::\n%s", hashEmail)
	// 5. Check OTP is available

	// 6. User spam

	// 1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrEmailAlreadyExists
	}

	// 2. New OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}
	fmt.Printf("OTP is :::%d \n", otp)
	// 3. Save OTP to Cache
	err := us.userAuthRepo.AddOtp(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		global.Logger.Error("save Otp to Redis error::", zap.Error(err))
		return response.ErrInvalidOtp
	}
	// 4. Send OTP to email
	err = sendto.SendTextEmailOtp([]string{email}, "quangtn0607@gmail.com", strconv.Itoa(otp))
	if err != nil {
		global.Logger.Error("Send email error", zap.Error(err))
		return response.ErrSendOtp
	}
	return response.ErrCodeSuccess
}
