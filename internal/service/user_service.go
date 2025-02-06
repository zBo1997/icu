package service

import (
	"errors"
	"icu/config"
	"icu/internal/model"
	"icu/internal/repository"
	"log"
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
func (s *UserService) Login(c *gin.Context)  (map[string]interface{}, error) {
	var reqUser model.User
	log.Println("login")
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		return nil, errors.New("invalid request")
	}

	// 检查用户是否存在
	dbUser, err := s.userRepo.UserExistByName(reqUser.Username)
	if err != nil || dbUser == nil {
		return nil, errors.New("user not found")
	}

	// 检查密码是否匹配
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(reqUser.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// 生成 JWT
	token, err := generateJWT(dbUser.Username)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

		// 返回用户信息（排除密码）和 Token
		userInfo := map[string]interface{}{
			"id":          dbUser.ID,
			"name":        dbUser.Name,
			"email":       dbUser.Email,
			"username":    dbUser.Username,
			"create_time": dbUser.CreateTime,
			"token":       token,
		}
		return userInfo, nil
}



// 注册处理函数
func  (s *UserService) Register(c *gin.Context) (*model.User, error) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return &user, errors.New("invalid request")
	}
	// 检查用户是否存在
	if user, _ := s.userRepo.UserExistByName(user.Username); user != nil {
		log.Println("user already exists")
		return user, errors.New("user already exists")
	}

	// 检查密码是否为空
	if user.Password == "" {
		log.Println("Password is required")
		return &user, errors.New("password is required")
	}

	// 检查密码是否为空
	if user.Username == "" {
		log.Println("username is required")
		return &user, errors.New("username is required")
	}
	return s.authRepo.SaveUser(user.Username,user.Password,user.Name,user.Email)
	
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