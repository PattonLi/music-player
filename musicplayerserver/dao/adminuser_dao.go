package dao

import (
	//"errors"
	"fmt"
	"music-player/musicplayerserver/model"
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
	DB.Find(&adminlist, "admin = ?", "true")
	return adminlist
}