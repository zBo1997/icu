package repository

import (
	"icu/config"
	"icu/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: config.GetDB()}
}

// GetUserByID 根据 ID 获取用户信息
func (r *UserRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UserExistByName 	根据用户名判断用户是否存在 如果不存在返回记录不存在错误
func (r *UserRepository) UserExistByName(name string) (*model.User, error) {
	var user model.User
	r.db.Where("name = ?",name).First(&user)
	return &user, nil
}

