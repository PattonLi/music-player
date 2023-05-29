package dao

import (
	"errors"
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
func (*UserDao) GetUserInfo(userID uint) (*model.UserInfo, error) {
	user := model.UserInfo{}
	err := DB.First(&user, "id = ?", userID).Error
	if err != nil {
		err = errors.New("查找不到用户信息！")
	}
	return &user, err
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

// 更新用户
func (*UserDao) UpdateUser(user *model.UserInfo) bool {
	result := DB.Save(user)
	if result.Error != nil {
		panic("Update error")
	} else {
		fmt.Println("Successfully update.")
		return true
	}
}

// 用户验证
func (*UserDao) UserCheck(u *model.UserInfo) (uint, string, error) {
	username := u.Username
	password := u.Password
	user := model.UserInfo{}
	err := DB.First(&user, "username = ? AND password = ?", username, password).Error
	if err != nil {
		err = errors.New("用户名不存在或密码错误！")
	}
	return user.ID, user.Admin, err
}

// 用户名验证
func (*UserDao) UsernameCheck(u *model.UserInfo) (uint, string, error) {
	username := u.Username
	user := model.UserInfo{}
	err := DB.First(&user, "username = ?", username).Error
	return user.ID, user.Admin, err
}
