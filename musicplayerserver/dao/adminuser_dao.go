package dao

import (
	//"errors"
	"errors"
	"fmt"
	"music-player/musicplayerserver/model"
	"gorm.io/gorm"
)

type AdminUserDao struct {
}

func NewAdminUserDao() *AdminUserDao {
	return &AdminUserDao{}
}

// 创建表，在第一次运行创建就行了，后面不用管他
func (*AdminUserDao) CreateAdminUserTable() {
	err := DB.AutoMigrate(&model.AdminUserInfo{})
	if err != nil {
		fmt.Println("Create AdminAUser Table failed: ", err)
	} else {
		fmt.Println("Successfully create AdminUser Table.")
	}
}

// 返回所有管理员信息
func (*AdminUserDao) GetAllAdminInfo() []model.AdminUserInfo {
	var adminlist []model.AdminUserInfo
	DB.Find(&adminlist)
	return adminlist
}

//根据adminname获取特定管理员信息
func (*AdminUserDao) GetAdminInfo(adminname string) (*model.AdminUserInfo,error) {
	adminuser := model.AdminUserInfo{}
	err := DB.Take(&adminuser,"adminname = ?", adminname).Error
	if err != nil{
		err = errors.New("查找不到该管理员！")
	}
	return &adminuser, err
}

//修改管理员信息
func (*AdminUserDao) ModifyAdminUser(adminuser *model.AdminUserInfo) error {
	result := DB.Save(adminuser)
	var err error = nil
	if result.Error != nil {
		err = errors.New("修改失败！")
	}
	return err
}

//添加管理员信息
func (*AdminUserDao) AddAdminUser(adminuser *model.AdminUserInfo) error {
	err := DB.Create(adminuser).Error
	return err
}

// 管理员用户名验证
func (*AdminUserDao) AdminUserNameCheck(au *model.AdminUserInfo) (int, error) {
	adminname := au.Adminname
	adminuser := model.AdminUserInfo{}
	err := DB.Take(&adminuser, "adminname = ?", adminname).Error
	if errors.Is(err, gorm.ErrRecordNotFound){
		err = errors.New("查找不到记录！")
	} else {
		err = nil
	}
	return adminuser.ID, err
}

//删除管理员信息
func (*AdminUserDao) DeleteAdminUser(adminID int) error {
	err := DB.Delete(&model.AdminUserInfo{}, adminID).Error
	return err
}