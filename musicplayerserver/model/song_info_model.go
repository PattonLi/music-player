package model

import (
	"gorm.io/gorm"
)

type SongInfo struct {
	gorm.Model
	Name      string
	Singer_id int
	Url       string
	Lyric     string
}

// 获取表名，gorm创建表时会自己获取这个表名
func (s *SongInfo) TableName() string {
	return "song"
}
