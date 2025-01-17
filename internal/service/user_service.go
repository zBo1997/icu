package service

import (
	"icu/config"
	"icu/internal/model"
	"icu/internal/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = config.GetKey("jwt:secret_key") // 用于签发 JWT 的密钥

// UserService 用于处理与用户相关的业务逻辑
type UserService struct {
	userRepo *repository.UserRepository
	authRepo *repository.AuthRepository
}

// NewUserService 创建 UserService 实例
func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(),
		authRepo: repository.NewAuthRepository(),
	}
}

// GetUser 根据用户 ID 获取用户信息
func (s *UserService) GetUser(id string) (*model.User, error) {
	return s.userRepo.GetUserByID(id)
}

// 登录处理函数
func (s *UserService) LoginHandler(c *gin.Context)  (string, error) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return "", nil
	}

	// 检查用户是否存在
	if _, err := s.userRepo.GetUserByID(user.Username); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return "", nil
	}
	
	// 检查密码是否匹配
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return "", nil
	}

	// 生成 JWT
	token, err := generateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return "", nil
	}
	return token, nil
}



// 注册处理函数
func  (s *UserService) RegisterHandler(c *gin.Context) (*model.User, error) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return &user, nil
	}

	// 检查密码是否为空
	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return &user, nil
	}

	// 检查密码是否为空
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return &user, nil
	}
	return s.authRepo.RegisterUser(user.Username,user.Password)
	
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
	return token.SignedString([]byte(jwtKey))
}