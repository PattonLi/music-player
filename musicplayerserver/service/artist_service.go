package service

import (
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
)

type ArtistService struct {
	artistdao *dao.AritistDao
}

func NewArtistService() *ArtistService {
	a := dao.AritistDao{}
	return &ArtistService{
		artistdao: &a,
	}
}

func (a *ArtistService) GetTenAritists() []model.ArtistInfo {
	artists := a.artistdao.GetTenArtist()
	return artists
}
