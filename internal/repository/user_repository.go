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

// GetUserByID 根据 ID 获取用户信息
func (r *UserRepository) GetUserByName(name string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("name = ?",name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

