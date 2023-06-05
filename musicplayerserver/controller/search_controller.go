package controller

import (
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"

	"github.com/gin-gonic/gin"
)

type SearchController struct {
	searchservice *service.SearchService
}

func (s *SearchController) SearchSuggestHandler(c *gin.Context) ([]model.AlbumInfo, []model.ArtistInfo, []model.SongInfo) {
	keywords := c.Query("keywords")
	album, artist, song := s.searchservice.GetSearchSuggest(keywords)
	return album, artist, song
}

func NewSearchController() *SearchController {
	ss := service.SearchService{}
	return &SearchController{
		searchservice: &ss,
	}
}
