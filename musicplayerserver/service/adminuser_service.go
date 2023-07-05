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

// 特定页所有管理员信息获取
func (aus *AdminUserService) AllAdminInfo(currentPage int, pageSize int) (int64, []model.AdminUserInfo) {
	totals, adminlist := aus.adminuserdao.GetAllAdminInfo(currentPage, pageSize)
	return totals, adminlist
}

//获取特定管理员信息
func (aus *AdminUserService) AdminUserInfo(adminname string) ([]model.AdminUserInfo,error) {
	adminlist, err := aus.adminuserdao.GetAdminInfo(adminname)
	return adminlist, err
}

//修改管理员信息
func (aus *AdminUserService) ModifyAdminUserInfo(adminuser *model.AdminUserInfo) error{
	err := aus.adminuserdao.ModifyAdminUser(adminuser)
	return err
}

//添加管理员信息
func (aus *AdminUserService) AddAdminUserInfo (adminuser *model.AdminUserInfo) (int64,int64,[]model.AdminUserInfo,error){
	_, err := aus.adminuserdao.AdminUserNameCheck(adminuser)
	var totals int64
	var currentPage int64
	var adminlist []model.AdminUserInfo
	if err != nil{
		totals, currentPage, adminlist, err =aus.adminuserdao.AddAdminUser(adminuser)
		err = nil
	} else {
		err = errors.New("管理员用户名已存在！")
	}
	return totals, currentPage, adminlist, err
}

//管理员登陆
func (aus *AdminUserService) AdminUserLogin(admin *model.AdminUserInfo) (int,error){
	adminID, err := aus.adminuserdao.AdminLoginCheck(admin)
	if err != nil{
		err = nil
		_, err = aus.adminuserdao.AdminUserNameCheck(admin)
		if err != nil{
			err = errors.New("账号输入错误！")
		} else {
			err = errors.New("密码输入错误！")
		}
	}
	return adminID, err
}

//根据ID获取单个管理员信息
func (aus *AdminUserService) AdminProfile(adminID int) (string, error) {
	adminname, err := aus.adminuserdao.GetAdminProfile(adminID)
	return adminname, err
}

//删除管理员信息
func (aus *AdminUserService) DeleteAdminUserInfo (adminID int) error {
	err := aus.adminuserdao.DeleteAdminUser(adminID)
	return err
}