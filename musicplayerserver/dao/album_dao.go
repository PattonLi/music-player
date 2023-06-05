package dao

import (
	"music-player/musicplayerserver/model"
)

type AlbumDao struct {
}

func (a *AlbumDao) GetTenAlbums() []model.AlbumInfo {
	var album []model.AlbumInfo
	var albums []model.AlbumInfo
	DB.Find(&album)
	albums = album[:10]
	return albums
}
