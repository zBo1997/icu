package model

import "time"

// Article 模型结构
type Article struct {
	ID            uint64   `gorm:"primarykey"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	UserId        uint64   `json:"userId"`
	AvatarUrl     string   `json:"avatarUrl"`
	Name          string   `json:"name"`
	TagNames      []string `gorm:"column:tag_names;serializer:json" json:"tagNames"`
	ImageKeys     []string `gorm:"column:image_keys;serializer:json" json:"imageKeys"`
	ImageCaptions []string `gorm:"column:image_captions;serializer:json" json:"imageCaptions"` // 新增
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
