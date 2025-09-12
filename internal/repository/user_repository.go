package repository

import (
	"icu/config"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 添加评论结构体
type User struct {
	gorm.Model
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	Signature string `json:"signature"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: config.GetDB()}
}

// repository/user_repository.go
func (r *UserRepository) FindMapByIds(ids []uint64) (map[uint64]User, error) {
	m := make(map[uint64]User, len(ids))
	if len(ids) == 0 {
		return m, nil
	}
	var users []User
	if err := r.db.Where("id IN ?", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	for _, u := range users {
		m[uint64(u.ID)] = u
	}
	return m, nil
}

// GetUserByID 根据 ID 获取用户信息
func (r *UserRepository) GetUserByID(id string) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UserExistByName 	根据用户名判断用户是否存在 如果不存在返回记录不存在错误
func (r *UserRepository) UserExistByName(name string) (*User, error) {
	var user User
	//按照名字寻找用户，如果没有则返回一个空的user
	if err := r.db.Where("username = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 注册
func (a *UserRepository) SaveUser(userName string, password string, name string, email string) (*User, error) {
	var user User
	// 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// 数据库保存用户信息
	user.Password = string(hashedPassword)
	user.Username = userName
	user.Name = name
	user.Email = email
	result := a.db.Create(&user)
	// 检查是否插入成功
	if result.Error != nil {
		return nil, result.Error
	}
	// 返回成功信息
	return &user, nil

}

// 修改用户头像
func (a *UserRepository) UpdateAvatar(user *User, imgKey string) (string, error) {
	//根据用户编号修改avatar字段未imgkey
	result := a.db.Model(&user).Update("avatar", imgKey).Where("id = ?", user.ID)
	if result.Error != nil {
		return "", result.Error
	}
	return imgKey, nil
}
