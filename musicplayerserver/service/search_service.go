package service

import (
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
)

type SearchService struct {
	sarchdao *dao.SearchDao
}

func (s *SearchService) GetSearchSuggest(keywords string) ([]model.AlbumInfo, []model.ArtistInfo, []model.SongInfo) {
	album, artist, song := s.sarchdao.GetKeywordInfo(keywords)
	return album, artist, song
}
