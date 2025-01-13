package controller

import (
	"fmt"
	"icu/internal/model"
	"icu/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret") // 用于签发 JWT 的密钥

// 用户结构体
type AuthController struct {
	userService *service.UserService
}

func NewAuthController() *AuthController{
	return &AuthController{
		userService: service.NewUserService(),
	}
}

// 用于保存用户信息的内存数据库（示例）
var users = map[string]string{}

// 登录处理函数
func (a *AuthController) LoginHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 检查用户是否存在
	storedPassword, exists := users[user.Username]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// 检查密码是否匹配
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// 生成 JWT
	token, err := generateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 返回 JWT
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// 注册处理函数
func  (a *AuthController) RegisterHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 检查密码是否为空
	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	// 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// 将用户信息保存到数据库（这里只是一个示例，实际中你需要使用数据库）
	users[user.Username] = string(hashedPassword)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
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
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}

	// 继续执行下一个请求
	c.Next()
}
// 私有用于生成 JWT 的函数
func generateJWT(username string) (string, error) {
	// 创建一个 JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	// 使用 HMAC 签名算法生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 返回 token 字符串
	return token.SignedString(jwtKey)
}