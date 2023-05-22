package model

import "gorm.io/gorm"

type SingerInfo struct {
	gorm.Model
	Name string
}

// 获取表名，gorm创建表时会自己获取这个表名
func (u *SingerInfo) TableName() string {
	return "singer"
}
