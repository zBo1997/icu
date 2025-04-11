package repository

import (
	"icu/config"
	"icu/internal/model"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{db: config.GetDB()}
}

//分页查询文章信息
func (a *ArticleRepository) FindAriticle(offset, limit int) ([]model.Article, int64, error) {
	var articles []model.Article
	var total int64

	err := a.db.Model(&model.Article{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = a.db.Table("articles").
		Select("articles.*, JSON_ARRAYAGG(tags.tag) as tag_names,users.avatar as avatar_url,users.name").
		Joins("LEFT JOIN article_tags ON article_tags.article_id = articles.id").
		Joins("LEFT JOIN tags ON tags.id = article_tags.tag_id").
		Joins("LEFT JOIN users ON users.id = articles.user_id").
		Group("articles.id").
		Order("articles.created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&articles).Error

	return articles, total, nil
}

//根据文章编号获取文章信息
func (a *ArticleRepository) GetArticle(articleId int) (model.Article, error) {
	var article model.Article
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