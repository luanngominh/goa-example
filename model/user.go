package model

type User struct {
	Model

	Email string `json:"email"`
	Fullname string `json:"fullname"`
	PasswordDigest string `json:"password_digest,omitempty"`
	VerifyCode string `json:"verify_code"`
	CodeExpireTime string `json:"code_expire_time"`
}