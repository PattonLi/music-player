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

func (s *SongInfo) TableName() string {
	return "song"
}
