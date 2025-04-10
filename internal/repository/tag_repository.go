package repository

import (
	"icu/config"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository() *TagRepository {
	return &TagRepository{db: config.GetDB()}
}
