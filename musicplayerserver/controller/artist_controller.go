package controller

import (
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"

	"github.com/gin-gonic/gin"
)

type ArtistController struct {
	artistservice *service.ArtistService
}

func NewArtistController() *ArtistController {
	a := service.ArtistService{}
	return &ArtistController{
		artistservice: &a,
	}
}

func (ac *ArtistController) GetTenArtistHandler(c *gin.Context) []model.ArtistInfo {
	artist := ac.artistservice.GetTenAritists()
	return artist
}
