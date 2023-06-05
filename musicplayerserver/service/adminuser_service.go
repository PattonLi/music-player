package service

import (
	"errors"
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
)

type AdminUserService struct {
	adminuserdao *dao.AdminUserDao
}

func NewAdminUserService() *AdminUserService {
	ausd := dao.NewAdminUserDao()
	return &AdminUserService{
		adminuserdao: ausd,
	}
}

// 所有管理员信息获取
func (aus *AdminUserService) AllAdminInfo() []model.AdminUserInfo {
	adminlist := aus.adminuserdao.GetAllAdminInfo()
	return adminlist
}

//获取特定管理员信息
func (aus *AdminUserService) AdminUserInfo(adminname string) (*model.AdminUserInfo,error) {
	adminuser, err := aus.adminuserdao.GetAdminInfo(adminname)
	return adminuser, err
}

//修改管理员信息
func (aus *AdminUserService) ModifyAdminUserInfo(adminuser *model.AdminUserInfo) error{
	err := aus.adminuserdao.ModifyAdminUser(adminuser)
	return err
}

//添加管理员信息
func (aus *AdminUserService) AddAdminUserInfo (adminuser *model.AdminUserInfo) error{
	_, err := aus.adminuserdao.AdminUserNameCheck(adminuser)
	if err != nil{
		aus.adminuserdao.AddAdminUser(adminuser)
		err = nil
	} else {
		err = errors.New("管理员用户名已存在！")
	}
	return err
}

//管理员登陆
func (aus *AdminUserService) AdminUserLogin(admin *model.AdminUserInfo) (uint,error){
	var err error = nil
	return 1, err
}

//删除管理员信息
func (aus *AdminUserService) DeleteAdminUserInfo (adminID int) error {
	err := aus.adminuserdao.DeleteAdminUser(adminID)
	return err
}