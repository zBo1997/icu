package controller

import (
	"encoding/base64"
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


// GetCaptchaHandler 获取验证码图片（Base64 + JSON 形式）
func (c *CaptchaController) GetCaptchaHandler(ctx *gin.Context) {
    captchaID ,captchaImage, err := c.service.GenerateCaptcha()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "message": "生成验证码失败",
        })
        return
    }
    // 把二进制转成 Base64
    imgB64 := base64.StdEncoding.EncodeToString(captchaImage)
    // 返回 JSON
    ctx.JSON(http.StatusOK, gin.H{"data": map[string]string{
        "captcha_id":    captchaID,
        "captcha_image": imgB64,
    }})
}
