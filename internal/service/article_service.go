package service

import (
	"icu/internal/model"
	"icu/internal/repository"
)

// UserService 用于处理与用户相关的业务逻辑
type ArticleService struct {
	articleRepo *repository.ArticleRepository
}

// NewUserService 创建 UserService 实例
func NewArticleService() *ArticleService {
	return &ArticleService{
		articleRepo: repository.NewArticleRepository(),
	}
}

// 分页获取文章列表
func (a *ArticleService) PageArticle(page, pageSize int) ([]model.Article, int64, error) {
	//计算偏移量 当前页数 * 每页数量
	offset := (page - 1) * pageSize
	return a.articleRepo.FindAriticle(offset, pageSize)
}

//根据文章编号获取文章信息
func (a *ArticleService) GetArticle(articleId int) (model.Article, error) {
	//查询文章信息
	return a.articleRepo.GetArticle(articleId)
}