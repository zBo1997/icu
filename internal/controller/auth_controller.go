package controller

import (
	"fmt"
	"icu/config"
	"icu/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// 用户结构体
type AuthController struct {
	userService *service.UserService
}

func NewAuthController() *AuthController{
	return &AuthController{
		userService: service.NewUserService(),
	}
}

// 登录处理函数
func (a *AuthController) LoginHandler(c *gin.Context) {
	token, err := a.userService.Login(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login user"})
		return
	}
	// 返回 JWT
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// 注册处理函数
func  (a *AuthController) RegisterHandler(c *gin.Context) {
	user, err := a.userService.Register(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
		
	}
	c.JSON(200, user)
}

// JWT 验证中间件
func  (a *AuthController) JwtMiddleware(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	if tokenStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(config.GetKey("jwt:secret_key")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}

	// 继续执行下一个请求
	c.Next()
}
