package model

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Name     string `json:"name" gorm:"column:name"`
	URL      string `json:"url" gorm:"column:url"`
	SingerID uint   `json:"singerid" gorm:"foreignKey:singer_id"`
	Singer   SingerInfo
}

// 获取表名，gorm创建表时会自己获取这个表名
func (u *Song) TableName() string {
	return "song"
}
