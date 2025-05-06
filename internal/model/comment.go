package model

import "gorm.io/gorm"

type Comment struct {
    gorm.Model
    ArticleID  int64  `json:"articleId"`          // 文章ID
    UserID     int64   `json:"userId"`             // 用户ID
    AvatarUrl  string `json:"avatarUrl"`          // 用户头像
    UserName   string `json:"name"`               // 用户名称
    Comment    string `json:"comment"`            // 评论内容
    ParentID   *int64 `json:"parentId,omitempty"` // 父评论ID（引用回复）
    ParentName string `json:"parentName"`         // 父评论用户名称（便于前端展示）
    LikesCount int    `json:"likesCount"`         // 点赞数量
}