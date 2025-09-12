package service

import (
	"icu/internal/repository"
)

// UserService 用于处理与用户相关的业务逻辑
type ArticleImageService struct {
	articleImageRepo *repository.ArticleImageRepository
}

// NewUserService 创建 UserService 实例
func NewArticleImageService() *ArticleImageService {
	return &ArticleImageService{
		articleImageRepo: repository.NewArticleImageRepository(),
	}
}
