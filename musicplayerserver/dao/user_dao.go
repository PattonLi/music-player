package dao

import (
	"errors"
	"fmt"
	"music-player/musicplayerserver/model"
	"math"
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



// 根据用户名查询用户信息
func (*UserDao) GetUserInfoByUsername(username string) ([]model.UserInfo, error) {
	var user []model.UserInfo
	err := DB.Find(&user, "username LIKE ?", username).Error
	if err != nil {
		err = errors.New("查找不到用户信息！")
	}
	return user, err
}

// 根据用户昵称查询用户信息
func (*UserDao) GetUserInfoByNickname(nickname string) ([]model.UserInfo, error){
	var user []model.UserInfo
	err := DB.Find(&user, "nickname LIKE ?", nickname).Error
	if err != nil {
		err = errors.New("查找不到用户信息！")
	}
	return user, err
}

// 添加用户
func (*UserDao) AddUser(user *model.UserInfo) bool {
	DB.Create(user)
	return true
}

// 修改用户
func (*UserDao) ModifyUser(user *model.UserInfo) error {
	result := DB.Save(user)
	var err error = nil
	if result.Error != nil {
		err = errors.New("修改失败!")
	}
	return err
}

// 用户验证
func (*UserDao) UserCheck(u *model.UserInfo) (uint,error) {
	username := u.Username
	password := u.Password
	user := model.UserInfo{}
	err := DB.First(&user, "username = ? AND password = ?", username, password).Error
	if err != nil {
		err = errors.New("用户名不存在或密码错误！")
	}
	return user.ID, err
}

// 用户名验证
func (*UserDao) UsernameCheck(u *model.UserInfo) (uint, error) {
	username := u.Username
	user := model.UserInfo{}
	err := DB.First(&user, "username = ?", username).Error
	return user.ID, err
}

// 返回特定页所有普通用户信息
func (*UserDao) GetAllUserInfo(page int, pagesize int) ([]model.UserInfo,int64) {
	var userlist []model.UserInfo
	var totalrecord int64
	offset := (page-1)*pagesize
	DB.Offset(offset).Limit(pagesize).Find(&userlist).Offset(-1).Limit(-1).Count(&totalrecord)
	totalPage := int64(math.Ceil(float64(totalrecord)/float64(pagesize)))
	return userlist,totalPage
}


