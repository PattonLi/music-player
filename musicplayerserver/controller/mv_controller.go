package controller

import (
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"
)

type MvController struct {
	mvservice *service.MvService
}

func NewMvController() *MvController {
	m := service.MvService{}
	return &MvController{
		mvservice: &m,
	}
}

func (mc *MvController) GetAllMvHandler() ([]model.MvInfo, error) {
	mvs, err := mc.mvservice.GetAllMv()
	return mvs, err
}
