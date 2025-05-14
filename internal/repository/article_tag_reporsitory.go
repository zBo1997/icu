package repository

import (
	"icu/config"

	"gorm.io/gorm"
)

type ArticleTag struct {
	gorm.Model
	ArticleId uint `json:"articleId"`
	TagId     uint `json:"tagId"`
}

type ArticleTagRepository struct {
	db *gorm.DB
}

func NewArticleTagRepository() *ArticleTagRepository {
	return &ArticleTagRepository{db: config.GetDB()}
}

// 保存文章和标签的关系
func (t *ArticleTagRepository) CreateArticleTag(articleId uint, param *ArticleTag) error {
	var tags []ArticleTag
	err := t.db.Table("article_tags").Where("tag_id = ?", param.TagId).Find(&tags).Error
	if err != nil {
		return err
	}
	//将标签插入到文章标签表中
	for _, tag := range tags {
		err = t.db.Table("article_tags").Create(&map[string]any{
			"article_id": articleId,
			"tag_id":     tag.ID,
		}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
