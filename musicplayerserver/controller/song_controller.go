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

// 获取歌词的控制器
func (sc *SongController) GetSongLyricHandler(c *gin.Context) string {
	type Name struct {
		Name string `json:"name"`
	}

	var name Name
	c.ShouldBindJSON(&name)
	n := name.Name
	lyric := sc.songservice.GetSongLyric(n)
	return lyric
}

// 获取歌曲详细信息的控制器
func (sc *SongController) GetSongDetailHandler(c *gin.Context) (string, string, error) {
	type parameter struct {
		Id string `json:"id"`
	}

	var par parameter
	c.ShouldBind(&par)
	id := par.Id
	songname, singer, err := sc.songservice.GetSongDetail(id)
	return songname, singer, err
}
