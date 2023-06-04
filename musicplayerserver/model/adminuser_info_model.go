package model

import (
	"gorm.io/gorm"
)

type AdminUserInfo struct {
	gorm.Model        // 创建对象时不用填充，gorm会自动填充
	ID			int    `gorm:"primaryKey;column:user_id"`
	Adminname   string `json:"adminname" gorm:"column:adminname"`
	Password   	string `json:"password" gorm:"column:password"`
}

// 获取表名，gorm创建表时会自己获取这个表名
func (a *AdminUserInfo) TableName() string {
	return "adminuser"
}