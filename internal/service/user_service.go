package service

import (
	"errors"
	"icu/config"
	"icu/internal/model"
	"icu/internal/repository"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = config.GetKey("jwt","secret_key") // 用于签发 JWT 的密钥

// UserService 用于处理与用户相关的业务逻辑
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService 创建 UserService 实例
func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(),
	}
}

// GetUser 根据用户 ID 获取用户信息
func (s *UserService) GetUser(id string) (*repository.User, error) {
	return s.userRepo.GetUserByID(id)
}

// 登录处理函数
func (s *UserService) Login(reqUser *model.User)  (map[string]interface{}, error) {
	
	// 检查用户是否存在
	dbUser, err := s.userRepo.UserExistByName(reqUser.Username)
	if err != nil || dbUser == nil {
		return nil, errors.New("未找到用户")
	}

	// 检查密码是否匹配
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(reqUser.Password))
	if err != nil {
		return nil, errors.New("密码错误")
	}

	// 生成 JWT
	token, err := generateJWT(dbUser.Username)
	if err != nil {
		return nil, errors.New("登录失败")
	}

		// 返回用户信息（排除密码）和 Token
		userInfo := map[string]interface{}{
			"id":          dbUser.ID,
			"name":        dbUser.Name,
			"email":       dbUser.Email,
			"username":    dbUser.Username,
			"create_time": dbUser.CreatedAt,
			"avatar":      dbUser.Avatar,
			"signature":   dbUser.Signature,
			"token":       token,
		}
		return userInfo, nil
}



// 注册处理函数
func  (s *UserService) Register(user *model.User) (*repository.User, error) {
	
	// 检查用户是否存在
	if user, _ := s.userRepo.UserExistByName(user.Username); user != nil {
		return nil, errors.New("用户已经存在")
	}

	// 检查密码是否为空
	if user.Password == "" {
		return nil, errors.New("请填写密码")
	}

	// 检查密码是否为空
	if user.Username == "" {
		return nil, errors.New("请填写用户名称")
	}
	return s.userRepo.SaveUser(user.Username,user.Password,user.Name,user.Email)
	
}

// 更新头像
func  (s *UserService) UpdateAvatar(c *gin.Context) (string , error) {
	var user *repository.User
    // 获取用户ID
    id := c.PostForm("userId")
    if id == "" {
        return "", errors.New("用户ID不能为空")
    }
    // 获取文件名
    fileName := c.PostForm("avatar")
    if fileName == "" {
        return "", errors.New("文件名不能为空")
    }
   	//获取用户信息
	user, err := s.userRepo.GetUserByID(id)
	//如果用户不存在//提示为空
	if err != nil{
		return "",errors.New("用户不存在")
	}
	return s.userRepo.UpdateAvatar(user,fileName)
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
	log.Println("签名"+jwtKey)
	// 返回 token 字符串
	return token.SignedString([]byte(jwtKey))
}