package dao

import (
	"fmt"
	"music-player/musicplayerserver/model"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

// 创建表，在第一次运行创建就行了，后面不用管他
func (*UserDao) CreateUserTable() {
	err := DB.AutoMigrate(&model.UserInfo{})
	if err != nil {
		fmt.Println("Create User Table failed: ", err)
	} else {
		fmt.Println("Successfully create User Table.")
	}
}

// 查询用户信息
func (*UserDao) GetUserInfo(id string) *model.UserInfo {
	user := model.UserInfo{}
	DB.First(&user, id)
	return &user
}

// 添加用户
func (*UserDao) AddUser(user *model.UserInfo) bool {
	result := DB.Create(user)
	if result.Error != nil {
		panic("Insert user error")
	} else {
		fmt.Println("Successfully insert user.")
		return true
	}
}
