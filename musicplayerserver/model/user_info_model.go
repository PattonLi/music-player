package model

import (

)

type UserInfo struct {
	ID   	   int    `json:"userId" gorm:"column:user_id;primaryKey;autoIncrement"`
	Username   string `json:"username" gorm:"column:username"`
	Gender     string `json:"gender" gorm:"column:gender"`
	Age        int    `json:"age" gorm:"column:age"`
	Email      string `json:"email" gorm:"column:email"`
	Password   string `json:"password" gorm:"column:password"`
	Nickname   string `json:"nickname" gorm:"column:nickname"`
	Phone      string `json:"phone" gorm:"column:phone"`
	Pic_url    string `json:"picUrl" gorm:"column:pic_url"`
}

// 获取表名，gorm创建表时会自己获取这个表名
func (u *UserInfo) TableName() string {
	return "user"
}
