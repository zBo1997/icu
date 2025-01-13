package model

// User 定义了用户的模型结构
type User struct {
	ID    string `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}
