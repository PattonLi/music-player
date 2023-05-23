package controller

import (
	"music-player/musicplayerserver/service"

	"github.com/gin-gonic/gin"
)

type SongController struct {
	songservice *service.SongService
}

func NewSongController() *SongController {
	ss := service.SongService{}
	return &SongController{
		songservice: &ss,
	}
}

func (sc *SongController) GetSongHandler(c *gin.Context) string {
	id := c.Query("id")
	lyric := sc.songservice.GetSong(id)
	return lyric
}
