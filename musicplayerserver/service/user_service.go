package service

import (
	"fmt"
	"musicplayerserver/dao"
	"musicplayerserver/model"
)

type UserService struct {
	userdao *dao.UserDao
}

func NewUserService() *UserService {
	usd := dao.NewUserDao()
	return &UserService{
		userdao: usd,
	}
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
