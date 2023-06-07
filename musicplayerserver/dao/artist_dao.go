package dao

import (
	"music-player/musicplayerserver/model"
)

type ArtistDao struct {
}

// 获取十个歌手信息
func (a *ArtistDao) GetTenArtist() []model.ArtistInfo {
	var artist []model.ArtistInfo
	var artists []model.ArtistInfo
	DB.Find(&artist)
	artists = artist[:10]
	return artists
}

// 根据歌手id获取歌手信息
func (a *ArtistDao) GetInfoById(id int) (model.ArtistInfo, error) {
	var artist model.ArtistInfo
	result := DB.First(&artist, id)
	return artist, result.Error
}

// 获取歌手描述
func (a *ArtistDao) GetProfile(id int) (string, error) {
	var artist model.ArtistInfo
	result := DB.First(&artist, id)
	return artist.Profile, result.Error
}
