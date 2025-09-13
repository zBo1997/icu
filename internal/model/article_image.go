package model

import "time"

// Article 模型结构
type ArticleWithImage struct {
	ID        uint64 `gorm:"primarykey"`
	ArticleID uint64 `json:"articleId"`
	Caption   string `json:"caption"`
	ImageKey  string `json:"imageKey"`
	SortOrder int    `json:"sortOrder"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
