package controller

import (
	"music-player/musicplayerserver/service"

	"github.com/gin-gonic/gin"
)

type SongController struct {
	songservice *service.SongService
}

func NewSongController() *SongController {
	ss := service.NewSongService()
	return &SongController{
		songservice: ss,
	}
}

func (sc *SongController) GetSongURLHandler(c *gin.Context) string {
	songname := c.PostForm("songname")
	url := sc.songservice.GetSongURL(songname)
	return url
}
