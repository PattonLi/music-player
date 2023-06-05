package dao

import (
	"music-player/musicplayerserver/model"
)

type AritistDao struct {
}

func (a *AritistDao) GetTenArtist() []model.ArtistInfo {
	var artist []model.ArtistInfo
	var artists []model.ArtistInfo
	DB.Find(&artist)
	artists = artist[:10]
	return artists
}
