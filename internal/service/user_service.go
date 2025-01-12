package service

import (
	"icu/internal/model"
	"icu/internal/repository"
)

// UserService 用于处理与用户相关的业务逻辑
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService 创建 UserService 实例
func NewUserService() *UserService {
	return &UserService{
		repo: repository.NewUserRepository(),
	}
}

// GetUser 根据用户 ID 获取用户信息
func (s *UserService) GetUser(id string) (*model.User, error) {
	return s.repo.GetUserByID(id)
}
