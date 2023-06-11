package service

import (
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
)

type MvService struct {
	mvdao *dao.MvDao
}

func (m *MvService) GetAllMv() ([]model.MvInfo, error) {
	mvs, err := m.mvdao.GetAllMv()
	return mvs, err
}
