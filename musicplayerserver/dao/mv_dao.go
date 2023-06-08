package dao

import (
	"music-player/musicplayerserver/model"
)

type MvDao struct {
}

func (m *MvDao) GetAllMv() ([]model.MvInfo, error) {
	var mv []model.MvInfo
	result := DB.Find(&mv)
	return mv, result.Error
}
