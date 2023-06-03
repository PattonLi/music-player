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




// 特定名称用户信息获取
func (us *UserService) UserInfo(name string) ([]model.UserInfo, error) {
	user1, err1 := us.userdao.GetUserInfoByUsername(name)
	user2, err2 := us.userdao.GetUserInfoByNickname(name)
	var err error = nil
	var userlist []model.UserInfo
	if(err1 != nil && err2 != nil){
		err = err2
	} else {
		userlist = append(user1,user2...)
	}
	return userlist, err
}

// 特定页所有普通用户信息获取
func (us *UserService) AllUserInfo(page int, pagesize int) ([]model.UserInfo, int64) {
	userlist,totalPage := us.userdao.GetAllUserInfo(page,pagesize)
	return userlist, totalPage
}

// 用户信息修改
func (us *UserService) ModifyUserInfo(user *model.UserInfo) {
	
}

// 用户登录
func (us *UserService) UserLogin(user *model.UserInfo) (uint, error) {
	//userID, admin, err := us.userdao.UserCheck(user)
	var err error = nil
	return 1, err
}

// 用户注册
func (us *UserService) UserRegister(user *model.UserInfo) error {
	_, err := us.userdao.UsernameCheck(user)
	if err != nil {
		us.userdao.AddUser(user)
		//user, _ = us.userdao.GetUserInfo(user.Username)
		err = nil
	} else {
		err = errors.New("用户名已被注册！")
	}
	return err
}
