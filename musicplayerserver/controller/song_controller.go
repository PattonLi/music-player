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

func (sc *SongController) GetSongLyricHandler(c *gin.Context) string {
	name := c.PostForm("name")
	lyric := sc.songservice.GetSong(name)
	return lyric
}
