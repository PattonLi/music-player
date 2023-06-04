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

//添加用户信息
func (us *UserService) AddUserInfo(user *model.UserInfo) error{
	var err error = nil
	return err
}

// 特定名称用户信息获取
func (us *UserService) UserInfo(name string) ([]model.UserInfo, error) {
	userlist, err := us.userdao.GetUserInfoByName(name)
	return userlist, err
}

// 特定页所有普通用户信息获取
func (us *UserService) AllUserInfo(page int, pagesize int) ([]model.UserInfo, int64) {
	userlist,totalPage := us.userdao.GetAllUserInfo(page,pagesize)
	return userlist, totalPage
}

// 用户信息修改
func (us *UserService) ModifyUserInfo(user *model.UserInfo) error{
	err := us.userdao.ModifyUser(user)
	return err
}

// 删除用户信息
func (us *UserService) DeleteUserInfo (userID int) error {
	err := us.userdao.DeleteUser(userID)
	return err
}

// 用户登录
func (us *UserService) UserLogin(user *model.UserInfo) (int, error) {
	userID, err := us.userdao.UserLoginCheck(user)
	if err != nil{
		err = nil
		_, err = us.userdao.UserPhoneCheck(user)
		if err != nil{
			err = errors.New("手机号输入错误！")
		} else {
			err = errors.New("密码输入错误！")
		}
	}
	return userID, err
}

// 用户注册
func (us *UserService) UserRegister(user *model.UserInfo) (int,error) {
	_, err := us.userdao.UserPhoneCheck(user)
	var userID int
	if err != nil {
		userID = us.userdao.AddUser(user)
		err = nil
	} else {
		err = errors.New("手机号已被注册！")
	}
	return userID,err
}

//根据ID获取单个用户信息
func (us *UserService) UserProfile(userID int) (*model.UserInfo, error) {
	user, err := us.userdao.GetUserProfile(userID)
	return user, err
}
