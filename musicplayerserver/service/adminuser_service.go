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
	adminuser := model.AdminUserInfo{}
	var err = errors.New("")
	return &adminuser, err
}

//管理员登陆
func (aus *AdminUserService) AdminUserLogin(admin *model.AdminUserInfo) (uint,error){
	var err error = nil
	return 1, err
}