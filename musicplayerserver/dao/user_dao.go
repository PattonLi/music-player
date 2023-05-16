package dao

import (
	"fmt"
	"music-player/musicplayerserver/model"
)

type UserDao struct {
}

// 创建表，在第一次运行创建就行了，后面不用管他
func (*UserDao) CreateUserTable() {
	err := DB.AutoMigrate(&model.UserInfo{})
	if err != nil {
		fmt.Println("Create Table failed: ", err)
	} else {
		fmt.Println("Successfully create Table.")
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
		panic("Insert error")
	} else {
		fmt.Println("Successfully insert.")
		return true
	}
}
