package service

import (
	"fmt"
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
)

type UserService struct {
	userdao *dao.UserDao
}

// 构造函数
func (*UserService) NewUserService() *UserService {
	usd := dao.UserDao{}
	us := UserService{
		userdao: &usd,
	}
	return &us
}

// 用户注册
func (us *UserService) UserRegister(user *model.UserInfo) {
	ok := us.userdao.AddUser(user)
	if ok {
		fmt.Println("Successful Register")
	} else {
		fmt.Println("Register Failed")
	}
}
