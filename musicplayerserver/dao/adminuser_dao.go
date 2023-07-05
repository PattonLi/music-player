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

// 返回特定页所有管理员信息
func (*AdminUserDao) GetAllAdminInfo(currentPage int, pageSize int) (int64, []model.AdminUserInfo) {
	var adminlist []model.AdminUserInfo
	var totals int64
	offset := (currentPage - 1) * pageSize
	fmt.Println(offset)
	fmt.Println(pageSize)
	DB.Offset(offset).Limit(pageSize).Find(&adminlist).Offset(-1).Limit(-1).Count(&totals)
	DB.Find(&adminlist)
	return totals, adminlist
}

//根据adminname获取特定管理员信息
func (*AdminUserDao) GetAdminInfo(adminname string) ([]model.AdminUserInfo,error) {
	var adminlist []model.AdminUserInfo
	err := DB.Find(&adminlist,"adminname LIKE ?", "%"+adminname+"%").Error
	if errors.Is(err, gorm.ErrRecordNotFound) || DB.RowsAffected == 0 {
		err = errors.New("查找不到管理员信息！")
	}
	return adminlist, err
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
func (*AdminUserDao) AddAdminUser(adminuser *model.AdminUserInfo) (int64, int64, []model.AdminUserInfo, error) {
	var adminlist []model.AdminUserInfo
	var totalrecord int64
	var offset int64
	var currentPage int64
	err := DB.Create(adminuser).Error
	DB.Table("admin_user").Count(&totalrecord)
	if(totalrecord % 10 == 0){
		offset = totalrecord - 10
		currentPage = totalrecord / 10
	} else {
		offset = totalrecord - (totalrecord % 10)
		currentPage = totalrecord / 10 + 1
	}
	DB.Offset(int(offset)).Limit(10).Find(&adminlist)
	return totalrecord, currentPage, adminlist, err
}

// 管理员登录验证
func (*AdminUserDao) AdminLoginCheck(au *model.AdminUserInfo) (int,error) {
	adminname := au.Adminname
	password := au.Password
	admin := model.AdminUserInfo{}
	err := DB.First(&admin, "adminname = ? AND password = ?", adminname, password).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("用户名不存在或密码错误！")
	}
	return admin.ID, err
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

//根据ID获取单个管理员信息
func (*AdminUserDao) GetAdminProfile(adminID int)(string, error) {
	admin := model.AdminUserInfo{}
	err := DB.First(&admin,"admin_id = ?", adminID).Error
	return admin.Adminname, err
}

//删除管理员信息
func (*AdminUserDao) DeleteAdminUser(adminID int) error {
	err := DB.Delete(&model.AdminUserInfo{}, adminID).Error
	return err
}