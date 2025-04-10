package model

import "gorm.io/gorm"

// Article 模型结构
type ArticleTags struct {
	gorm.Model
	ArticleId string `json:"title"`
	TagId     string `json:"tagId"`
}
