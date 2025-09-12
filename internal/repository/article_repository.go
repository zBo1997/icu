package repository

import (
	"icu/config"
	"icu/internal/model"

	"gorm.io/gorm"
)

// Article 模型结构
type Article struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  uint64 `json:"userId"`
}

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{db: config.GetDB()}
}

// 分页查询文章信息
func (a *ArticleRepository) FindArticle(offset, limit int) ([]Article, int64, error) {
	var articles []Article
	var total int64

	err := a.db.Model(&Article{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = a.db.Table("articles").
		Select("articles.*").
		Order("articles.created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&articles).Error

	return articles, total, nil
}

// 根据文章编号获取文章信息
func (a *ArticleRepository) GetArticle(articleId int) (model.ArticleWithImage, error) {
	var article model.ArticleWithImage
	// 查询文章信息
	err := a.db.Table("articles").
		Select("articles.*, JSON_ARRAYAGG(tags.tag) as tag_names, users.avatar as avatar_url, users.name").
		Joins("LEFT JOIN article_tags ON article_tags.article_id = articles.id").
		Joins("LEFT JOIN tags ON tags.id = article_tags.tag_id").
		Joins("LEFT JOIN users ON users.id = articles.user_id").
		Where("articles.id = ?", articleId).
		Group("articles.id").
		First(&article).Error

	if err != nil {
		return article, err
	}

	return article, nil
}

// 发布文章
func (a *ArticleRepository) CreateArticle(article *Article) (uint64, error) {
	err := a.db.Create(article).Error
	if err != nil {
		return 0, err
	}
	return uint64(article.ID), nil
}
