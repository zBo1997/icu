package model

import "gorm.io/gorm"

// User 定义了用户的模型结构
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Avatar  string `json:"avatar"`
	Email string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Signature string `json:"signature"`
	CaptchaID string `json:"captchaID"`
	CaptchaCode string `json:"captchaCode"`
}
