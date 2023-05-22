package controller

import (
	"music-player/musicplayerserver/service"

	"github.com/gin-gonic/gin"
)

var songservice = service.SongService{}

func SongControllerInit() {
	songservice.NewSongService()
}

func GetSongHandler(c *gin.Context) string {
	id := c.Query("id")
	lyric := songservice.GetSong(id)
	return lyric
}
