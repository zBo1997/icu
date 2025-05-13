package repository

import (
	"icu/config"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Tag    string `json:"title"`
	UserId string `json:"userId"`
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

