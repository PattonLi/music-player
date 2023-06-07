package service

import (
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
)

type ArtistService struct {
	artistdao *dao.ArtistDao
}

func NewArtistService() *ArtistService {
	a := dao.ArtistDao{}
	return &ArtistService{
		artistdao: &a,
	}
}

// 获取十个歌手信息
func (a *ArtistService) GetTenAritists() []model.ArtistInfo {
	artists := a.artistdao.GetTenArtist()
	return artists
}

// 根据歌手id获取歌手信息
func (a *ArtistService) GetArtistDetail(id int) (model.ArtistInfo, error) {
	artist, err := a.artistdao.GetInfoById(id)
	return artist, err
}

// 获得歌手描述
func (a *ArtistService) GetArtistDescribe(id int) (string, error) {
	profile, err := a.artistdao.GetProfile(id)
	return profile, err
}
