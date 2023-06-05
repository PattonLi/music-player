package service

import (
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
)

type AlbumService struct {
	albumdao *dao.AlbumDao
}

func NewAlbumService() *AlbumService {
	a := dao.AlbumDao{}
	al := AlbumService{
		albumdao: &a,
	}
	return &al
}

func (al *AlbumService) GetTenAlbums() []model.AlbumInfo {
	albums := al.albumdao.GetTenAlbums()
	return albums
}
