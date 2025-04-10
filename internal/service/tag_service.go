package service

import (
	"icu/internal/model"
	"icu/internal/repository"
)

// UserService 用于处理与用户相关的业务逻辑
type TagService struct {
	TagRepository *repository.TagRepository
}

// NewUserService 创建 UserService 实例
func NewTagService() *TagService {
	return &TagService{
		TagRepository: repository.NewTagRepository(),
	}
}

// GetUser 根据用户 ID 获取用户信息
func (s *TagService) getTagList(id string) (*model.Tag, error) {
	return nil, nil
}
