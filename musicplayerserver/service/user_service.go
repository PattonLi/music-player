package service

import (
	"errors"
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
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
func (us *UserService) UserRegister(user *model.UserInfo) error {
	_, _, err := us.userdao.UsernameCheck(user)
	if err != nil {
		us.userdao.AddUser(user)
		user, _ = us.userdao.GetUserInfo(user.Username)
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

// 单个用户信息获取
func (us *UserService) UserInfo(username string) (*model.UserInfo, error) {
	user, err := us.userdao.GetUserInfo(username)
	return user, err
}

// 所有普通用户信息获取
func (us *UserService) AllUserInfo() []model.UserInfo {
	userlist := us.userdao.GetAllUserInfo()
	return userlist
}

// 所有管理员信息获取
func (us *UserService) AllAdminInfo() []model.UserInfo {
	adminlist := us.userdao.GetAllAdminInfo()
	return adminlist
}

// 用户信息修改
func (us *UserService) UserEdit() {

}
