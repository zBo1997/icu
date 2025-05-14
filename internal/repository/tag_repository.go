package repository

import (
	"icu/config"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Tag    string `json:"title"`
	UserId int64 `json:"userId"`
}

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository() *TagRepository {
	return &TagRepository{db: config.GetDB()}
}

func (t *TagRepository) GetTagList() ([]Tag, error) {
	var tagList []Tag
	err := t.db.Find(&tagList).Error
	if err != nil {
		return nil, err
	}
	return tagList, nil
}

func (t *TagRepository) CreateArticleTag(articleId uint, param *Tag) error {
	var tags []Tag
	err := t.db.Table("tags").Where("tag IN ?", param).Find(&tags).Error
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