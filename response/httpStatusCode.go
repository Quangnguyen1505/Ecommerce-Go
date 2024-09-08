package response

const (
	ErrCodeSuccess         = 20001
	ErrCodeParamInvalid    = 20003
	ErrTokenHeadersInvalid = 30002
	//check email
	ErrEmailAlreadyExists = 50002
)

var msg = map[int]string{
	ErrCodeSuccess:         "success",
	ErrCodeParamInvalid:    "Email invalid",
	ErrTokenHeadersInvalid: "Token headers invalid",
	ErrEmailAlreadyExists:  "Email Already exists",
}
