package model

import "gorm.io/gorm"

// Tag 模型结构
type Tag struct {
	gorm.Model
	Tag    string `json:"title"`
	UserId string `json:"userId"`
}
