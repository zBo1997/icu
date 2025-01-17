package repository

import (
	"icu/config"
	"icu/internal/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{db: config.GetDB()}
}

// 注册
func  (a *AuthRepository) RegisterUser(userName string,password string)(*model.User, error) {
	var user model.User
	// 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// 数据库保存用户信息
	user.Password = string(hashedPassword)
	if err != nil {
		return nil, err
	}

	// 返回成功信息
	return  &user, nil

}

// GetUserByID 根据 ID 获取用户信息
func (a *AuthRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User
	if err := a.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}