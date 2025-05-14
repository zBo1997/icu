package service

import (
	"icu/config"
	"icu/internal/model"
	"icu/internal/repository"
)

// UserService 用于处理与用户相关的业务逻辑
type ArticleService struct {
	articleRepo *repository.ArticleRepository
	articleTagRepository *repository.ArticleTagRepository
}

// NewUserService 创建 UserService 实例
func NewArticleService() *ArticleService {
	return &ArticleService{
		articleRepo: repository.NewArticleRepository(),
		articleTagRepository: repository.NewArticleTagRepository(),
	}
}

// 分页获取文章列表
func (a *ArticleService) PageArticle(page, pageSize int) ([]model.Article, int64, error) {
	//计算偏移量 当前页数 * 每页数量
	offset := (page - 1) * pageSize
	return a.articleRepo.FindAriticle(offset, pageSize)
}

//根据文章编号获取文章信息
func (a *ArticleService) GetArticle(articleId int) (model.ArticleWithImage, error) {
	//查询文章信息
	return a.articleRepo.GetArticle(articleId)
}

//发布文章
func (a *ArticleService) PublishArticle(article *model.ArticlePublish, userId int64) (uint, error) {

	// 开启事务
	tx := config.GetDB().Begin()
	//创建发布文章对象
	articleModel := repository.Article{
		Title:    article.Title,
		Content:  article.Content,
		ImageKey: article.ImageKey,
		UserId:   userId,
	}
	//保存文章
	articleId, err := a.articleRepo.CreateArticle(&articleModel)
	if err != nil {
		tx.Rollback() // 回滚事务
		return 0, err
	}

	tag := &repository.ArticleTag{
		ArticleId:    articleId,
		TagId: article.Tags,
	}

	//保存文章标签关系
	err = a.articleTagRepository.CreateArticleTag(articleId, tag)
	if err != nil {
		tx.Rollback() // 回滚事务
		return 0, err
	}

	// 提交事务
    if err := tx.Commit().Error; err != nil {
        return 0, err
    }
	return articleId, nil
}