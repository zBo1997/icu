package service

import (
	"bytes"

	"github.com/dchest/captcha"
)

type CaptchaService struct{}

func NewCaptchaService() *CaptchaService {
	return &CaptchaService{}
}

// GenerateCaptcha 生成验证码，返回验证码 ID 以及对应的 PNG 图片字节
func (s *CaptchaService) GenerateCaptcha() (captchaID string, imgBytes []byte, err error) {
	// 生成一个默认长度（6 位）的验证码 ID
	captchaID = captcha.New()

	// 把对应 ID 的验证码写成 PNG，尺寸 240×80
	var buf bytes.Buffer
	if err = captcha.WriteImage(&buf, captchaID, 240, 80); err != nil {
		return "", nil, err
	}

	// 返回 ID 和图片数据
	return captchaID, buf.Bytes(), nil
}

// VerifyCaptcha 根据 ID 和用户提交的值做校验，返回 true 表示通过
func (s *CaptchaService) VerifyCaptcha(captchaID, userValue string) bool {
	// VerifyString 内部会一次性删除该 ID，防止重放
	return captcha.VerifyString(captchaID, userValue)
}