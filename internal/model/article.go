package model

import "gorm.io/gorm"

// Article 模型结构
type Article struct {
	gorm.Model
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	UserId   string   `json:"userId"`
	TagNames []string `gorm:"column:tag_names;serializer:json"`
}
