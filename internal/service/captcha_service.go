package service

import (
	"bytes"

	"github.com/dchest/captcha"
)

type CaptchaService struct{}

func NewCaptchaService() *CaptchaService {
	return &CaptchaService{}
}

// GenerateCaptcha 生成验证码并返回验证码 ID 和图片数据
func (s *CaptchaService) GenerateCaptcha() (string, []byte, error) {
	captchaID := captcha.New()
	var buffer bytes.Buffer
	if err := captcha.WriteImage(&buffer, captchaID, 240, 80); err != nil {
		return "", nil, err
	}
	return captchaID, buffer.Bytes(), nil
}
