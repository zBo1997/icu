package controller

import (
	"net/http"

	"icu/internal/service"

	"github.com/gin-gonic/gin"
)

// CaptchaController 用于处理验证码相关的请求
type CaptchaController struct {
	service *service.CaptchaService
}

// NewCaptchaController 创建一个新的 CaptchaController 实例
func NewCaptchaController() *CaptchaController {
	return &CaptchaController{
		service: service.NewCaptchaService(),
	}
}

// GetCaptchaHandler 获取验证码图片
func (c *CaptchaController) GetCaptchaHandler(ctx *gin.Context) {
	captchaID, captchaImage, err := c.service.GenerateCaptcha()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate captcha",
		})
		return
	}

	// 将验证码图片以二进制流的形式返回
	ctx.Header("Content-Type", "image/png")
	ctx.Header("Captcha-ID", captchaID) // 可选：返回验证码 ID 供前端验证时使用
	ctx.Data(http.StatusOK, "image/png", captchaImage)
}
