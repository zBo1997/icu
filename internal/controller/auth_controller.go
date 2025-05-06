package controller

import (
	"encoding/json"
	"icu/internal/model"
	"icu/internal/service"
	"icu/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户结构体
type AuthController struct {
	userService *service.UserService
	captchaService *service.CaptchaService
}

func NewAuthController() *AuthController{
	return &AuthController{
		userService: service.NewUserService(),
		captchaService: service.NewCaptchaService(),
	}
}

// 登录处理函数
func (a *AuthController) LoginHandler(c *gin.Context) {
	var reqUser model.User
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "非法请求"})
		return 
	}
	//打印请求参数 以JSON格式
	result,_ := json.Marshal(reqUser)

	log.Println("请求参数", string(result))

	verify := a.captchaService.VerifyCaptcha(reqUser.CaptchaID, reqUser.CaptchaCode)

	if (!verify) {
		c.JSON(http.StatusOK, gin.H{"data": map[string]string{"error": "验证码错误"}})
		return
	}

	userInfo, err := a.userService.Login(&reqUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": map[string]string{"error": err.Error()}})
		return
	}
	// 返回 JWT
	c.JSON(http.StatusOK, gin.H{"data": userInfo})
}

// 注册处理函数
func  (a *AuthController) RegisterHandler(c *gin.Context) {
	var reqUser model.User
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "非法请求"})
		return 
	}

	//打印请求参数 以JSON格式
	result,_ := json.Marshal(reqUser)

	log.Println("请求参数", string(result))

	verify := a.captchaService.VerifyCaptcha(reqUser.CaptchaID, reqUser.CaptchaCode)

	if (!verify) {
		c.JSON(http.StatusOK, gin.H{"data": map[string]string{"error": "验证码错误"}})
		return
	}

	user, err := a.userService.Register(&reqUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": map[string]string{"error": err.Error()}})
		return
		
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// JWT 验证中间件
func (a *AuthController) JwtMiddleware(c *gin.Context) {
    tokenStr := c.GetHeader("Authorization")
    if tokenStr == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"data": map[string]string{"error": "未登录,请登陆后再试"}})
        c.Abort()
        return
    }

    // 使用工具类解析 userId
    userId, err := utils.ParseUserIDFromToken(tokenStr)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"data": map[string]string{"error": err.Error()}})
        c.Abort()
        return
    }

    // 将 userId 存储到上下文中 这样后续的处理函数可以直接获取
    c.Set("userId", userId)

    // 继续执行下一个请求
    c.Next()
}
