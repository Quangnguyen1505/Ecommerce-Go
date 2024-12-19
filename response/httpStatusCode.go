package response

const (
	ErrCodeSuccess         = 20001
	ErrCodeParamInvalid    = 20003
	ErrTokenHeadersInvalid = 30002
	ErrCodeRemoveSuccess   = 2000302
	ErrInvalidOtp          = 30003
	ErrSendOtp             = 30004
	//check email
	ErrEmailAlreadyExists = 50002

	ErrCodeUserHasExists = 60002
	ErrCodeOtpNotExists  = 60003
	ErrCodeOtpNotVerify  = 60004

	// authentication
	ErrCodeUserNotRegister = 60006
	ErrCodeAuthenError     = 60005

	//Redis
	ErrCodeRedisSetFailed = 70001
	ErrCodeRedisGetFailed = 70002

	// Two Factor Authentication
	ErrTwoFactorAuthSetUpFailed = 80002
)

var msg = map[int]string{
	ErrCodeSuccess:              "success",
	ErrCodeParamInvalid:         "Email invalid",
	ErrTokenHeadersInvalid:      "Token headers invalid",
	ErrEmailAlreadyExists:       "Email Already exists",
	ErrInvalidOtp:               "Error Otp",
	ErrSendOtp:                  "Send Otp error",
	ErrCodeRemoveSuccess:        "test successful",
	ErrCodeUserHasExists:        "key user already exists",
	ErrCodeOtpNotExists:         "OTP exists but not registered",
	ErrCodeOtpNotVerify:         "OTP exist but not verifyed",
	ErrCodeAuthenError:          "Authentication error",
	ErrCodeUserNotRegister:      "User don't exists",
	ErrCodeRedisSetFailed:       "Redis Set in RAM failed",
	ErrCodeRedisGetFailed:       "Redis Get in RAM failed",
	ErrTwoFactorAuthSetUpFailed: "Two factor authentication Error",
}
