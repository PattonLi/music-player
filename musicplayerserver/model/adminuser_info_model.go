package model

import (

)

type AdminUserInfo struct {
	ID			int    `json:"adminId" gorm:"primaryKey;column:admin_id"`
	Adminname   string `json:"adminName" gorm:"column:adminname"`
	Password   	string `json:"password" gorm:"column:password"`
}

// 获取表名，gorm创建表时会自己获取这个表名
func (a *AdminUserInfo) TableName() string {
	return "adminuser"
}