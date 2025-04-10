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
func (a *ArticleRepository) FindAriticle(offset, limit int) ([]model.Article, int64, error) {
	var articles []model.Article
	var total int64

	err := a.db.Model(&model.Article{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = a.db.Table("articles").
		Select("articles.*, JSON_ARRAYAGG(tags.tag) as tag_names").
		Joins("LEFT JOIN article_tags ON article_tags.article_id = articles.id").
		Joins("LEFT JOIN tags ON tags.id = article_tags.tag_id").
		Group("articles.id").
		Offset(offset).
		Limit(limit).
		Find(&articles).Error

	return articles, total, nil
}
