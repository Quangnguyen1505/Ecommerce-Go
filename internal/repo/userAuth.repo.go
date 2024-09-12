package repo

import (
	"fmt"
	"time"

	"github.com/ntquang/ecommerce/global"
)

type IUserAuthRepository interface {
	AddOtp(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

func (r *userAuthRepository) AddOtp(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("usr:%s:otp", email)
	return global.Redis.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
