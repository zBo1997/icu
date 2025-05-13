package service

import (
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
func (s *TagService) GetTagList() ([]repository.Tag, error) {
	// 调用 TagRepository 的 GetTagList 方法
	tagList, err := s.TagRepository.GetTagList()
	if err != nil {
		return nil, err
	}
	// 返回用户信息
	return tagList, nil
}
