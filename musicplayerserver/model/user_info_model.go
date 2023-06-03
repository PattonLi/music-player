package model

import (
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model        // 创建对象时不用填充，gorm会自动填充
	ID   	   uint   `gorm:"primaryKey;column:user_id"`
	Username   string `json:"username" gorm:"column:username"`
	Gender     string `json:"gender" gorm:"column:gender"`
	Age        string `json:"age" gorm:"column:age"`
	Email      string `json:"email" gorm:"column:email"`
	Password   string `json:"password" gorm:"column:password"`
	Nickname   string `json:"nickname" gorm:"column:nickname"`
	Phone      string `json:"phone" gorm:"column:phone"`
	Pic_url    string `json:"pic_url" gorm:"column:pic_url"`
}

// 获取表名，gorm创建表时会自己获取这个表名
func (u *UserInfo) TableName() string {
	return "user"
}
