package dao

import (
	"music-player/musicplayerserver/model"
)

type SearchDao struct {
}

func (s *SearchDao) GetKeywordInfo(keywords string) ([]model.AlbumInfo, []model.ArtistInfo, []model.SongInfo) {
	var album []model.AlbumInfo
	var artist []model.ArtistInfo
	var song []model.SongInfo

	keywords = "%" + keywords + "%"

	DB.Where("name  LIKE ?", keywords).Find(&album)
	DB.Where("name  LIKE ?", keywords).Find(&artist)
	DB.Where("name  LIKE ?", keywords).Find(&song)

	return album, artist, song
}
