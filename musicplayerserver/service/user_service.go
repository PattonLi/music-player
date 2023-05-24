package service

import (
	"errors"
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
func (us *UserService) UserRegister(user *model.UserInfo) error {
	_, _, err := us.userdao.UsernameCheck(user)
	if err != nil {
		ok := us.userdao.AddUser(user)
		if ok {
			fmt.Println("Successful Register")
		} else {
			fmt.Println("Register Failed")
		}
		err = nil
	} else {
		err = errors.New("用户名已被注册！")
	}
	return err
}

// 用户登录
func (us *UserService) UserLogin(user *model.UserInfo) (uint, string, error) {
	userID, admin, err := us.userdao.UserCheck(user)
	return userID, admin, err
}

// 用户信息获取
func (us *UserService) UserInfo(userID uint) (*model.UserInfo, error) {
	user, err := us.userdao.GetUserInfo(userID)
	return user, err
}

// 用户登出
func (us *UserService) UserLogout() {

}

// 用户信息修改
func (us *UserService) UserEdit() {

}
