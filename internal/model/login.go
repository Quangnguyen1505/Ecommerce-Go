package model

type RegisterInput struct {
	VerifyKey     string `json:"verify_key"`
	VerifyType    int    `json:"verify_type"`
	VerifyPurpose string `json:"verify_purpose"`
}

type VerifyInput struct {
	VerifyCode string `json:"verify_code"`
	VerifyKey  string `json:"verify_key"`
}

type VerifyOTPOutput struct {
	Token   string `json:"token"`
	UserId  string `json:"user_id"`
	Message string `json:"message"`
}
