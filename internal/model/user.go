package model

// User 定义了用户的模型结构
type User struct {
	ID    int32 `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Avatar  string `json:"avatar"`
	Email string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Signature string `json:"signature"`
	CreateTime int64 `json:"create_time" gorm:"-"`
}
