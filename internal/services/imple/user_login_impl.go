package imple

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ntquang/ecommerce/global"
	consts "github.com/ntquang/ecommerce/internal/const"
	"github.com/ntquang/ecommerce/internal/database"
	"github.com/ntquang/ecommerce/internal/model"
	"github.com/ntquang/ecommerce/internal/utlis"
	"github.com/ntquang/ecommerce/internal/utlis/crypto"
	"github.com/ntquang/ecommerce/internal/utlis/random"
	"github.com/ntquang/ecommerce/internal/utlis/sendto"
	"github.com/ntquang/ecommerce/response"
	"github.com/redis/go-redis/v9"
)

type sUserLogin struct {
	// implement the IUserLogin interface here
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}

// implement the IUserLogin interface here

func (s *sUserLogin) Login(ctx context.Context) error {
	fmt.Println("login already work")
	return nil
}

func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (resultCode int, err error) {
	fmt.Printf("VerifyKey: %s\n", in.VerifyKey)
	fmt.Printf("VerifyType: %d\n", in.VerifyType)
	// 0. hash VerifyKey
	hashKey := crypto.GetHash(in.VerifyKey)
	fmt.Printf("hashKey: %s\n", hashKey)
	// 1 check user exists
	userFound, err := s.r.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeUserHasExists, err
	}

	fmt.Println("USER::", userFound)

	if userFound > 0 {
		return response.ErrCodeUserHasExists, fmt.Errorf("user has already registered")
	}

	//3 create OTP
	userKey := utlis.GetUserKey(hashKey)
	otpFound, err := global.Redis.Get(ctx, userKey).Result()

	switch {
	case err == redis.Nil:
		fmt.Println("Key doesn't exists!")
	case err != nil:
		fmt.Println("get failed ::", err)
		return response.ErrInvalidOtp, err
	case otpFound != "":
		return response.ErrCodeOtpNotExists, fmt.Errorf("")
	}

	//4 generate OTP
	otpNew := random.GenerateSixDigitOtp()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Printf("otpNew: %s\n", otpNew)

	//5. save OTP in redos
	err = global.Redis.Set(ctx, userKey, strconv.Itoa(otpNew), time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrInvalidOtp, err
	}

	switch in.VerifyType {
	case consts.EMAIL:
		err = sendto.SendTextEmailOtp([]string{in.VerifyKey}, consts.EMAIL_SEND, strconv.Itoa(otpNew))
		if err != nil {
			return response.ErrSendOtp, err
		}
		// 7. save OTP to Postgresql
		result, err := s.r.InsertOtpVerify(ctx, database.InsertOtpVerifyParams{
			VerifyOtp:     strconv.Itoa(otpNew),
			VerifyKey:     in.VerifyKey,
			VerifyType:    pgtype.Int4{Int32: 1, Valid: true},
			VerifyKeyHash: hashKey,
		})
		if err != nil {
			return response.ErrSendOtp, err
		}
		//8. getlastId when have function work need last ID
		// lastIdVerifyUser, err := result.
		fmt.Println("lastId ::", result)
		return response.ErrCodeSuccess, nil
	case consts.MOBILE:
		return response.ErrCodeSuccess, nil
	}
	return response.ErrCodeSuccess, nil
}

func (s *sUserLogin) VerifyOtp(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error) {
	// 1 hash email
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	// 2 get otp
	userKey := utlis.GetUserKey(hashKey)
	otpFound, err := global.Redis.Get(ctx, userKey).Result()
	if err != nil {
		return out, err
	}

	if in.VerifyCode != otpFound {
		// if failed count as three throw error
		// ... login
		return out, fmt.Errorf("Otp not match")
	}

	infoOTP, err := s.r.GetInfoOtp(ctx, hashKey)
	if err != nil {
		return out, err
	}

	// 3. update status
	err = s.r.UpdateUserValificationStatus(ctx, hashKey)
	if err != nil {
		return out, fmt.Errorf("Update failed!!")
	}

	out.Token = infoOTP.VerifyKeyHash
	out.Message = "Verify OTP successfull"

	return out, err
}

func (s *sUserLogin) UpdatePaswordRegister(ctx context.Context) error {
	return nil
}
