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
)

var msg = map[int]string{
	ErrCodeSuccess:         "success",
	ErrCodeParamInvalid:    "Email invalid",
	ErrTokenHeadersInvalid: "Token headers invalid",
	ErrEmailAlreadyExists:  "Email Already exists",
	ErrInvalidOtp:          "Error Otp",
	ErrSendOtp:             "Send Otp error",
	ErrCodeRemoveSuccess:   "test successful",
	ErrCodeUserHasExists:   "key user already exists",
	ErrCodeOtpNotExists:    "OTP exists but not registered",
}
