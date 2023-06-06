package dao

import (
	"music-player/musicplayerserver/model"
)

type AlbumDao struct {
}

// 从数据库中获取十张专辑
func (a *AlbumDao) GetTenAlbums() []model.AlbumInfo {
	var album []model.AlbumInfo
	var albums []model.AlbumInfo
	DB.Find(&album)
	albums = album[:10]
	return albums
}

// 根据专辑id获取专辑
func (a *AlbumDao) GetAlbumById(id int) (model.AlbumInfo, error) {
	var album model.AlbumInfo
	result := DB.First(&album, id)

	return album, result.Error
}

// 根据歌手id获取歌手所有专辑
func (a *AlbumDao) GetAlbumByArtistid(artist_id int) ([]model.AlbumInfo, error) {
	var album []model.AlbumInfo
	result := DB.Where("artist_id = ?", artist_id).Find(&album)
	return album, result.Error
}
