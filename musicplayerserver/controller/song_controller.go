package controller

import (
	"fmt"
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"
	"strconv"

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
func (sc *SongController) GetSongLyricHandler(c *gin.Context) (string, error) {
	/*type Name struct {
		Name string `json:"name"`
	}

	var name Name
	c.ShouldBindJSON(&name)
	n := name.Name*/
	id := c.Query("id")
	lyric, err := sc.songservice.GetSongLyric(id)
	return lyric, err
}

/*
// 获取歌曲详细信息的控制器
func (sc *SongController) GetSongDetailHandler(c *gin.Context) (string, string, error) {
	/*type parameter struct {
		Id string `json:"id"`
	}

	var par parameter
	c.ShouldBind(&par)
	id := par.Id
	id := c.Query("id")
	songname, singer, err := sc.songservice.GetSongDetail(id)
	return songname, singer, err
}*/

// 添加歌曲的控制器
func (sc *SongController) AddSongHandler(c *gin.Context) bool {
	var song model.SongInfo
	c.ShouldBind(&song)
	result := sc.songservice.AddSong(song)
	return result
}

// 获取十首歌曲
func (sc *SongController) GetTenSongsHandler(c *gin.Context) []model.SongInfo {
	songs := sc.songservice.GetTenSongs()
	return songs
}

// 获取专辑的所有歌曲
func (sc *SongController) GetAlbumSongsHandler(c *gin.Context) ([]model.SongInfo, error) {
	albumid := c.Query("albumId")
	albumId, err0 := strconv.Atoi(albumid)
	if err0 != nil {
		fmt.Println("字符串转换错误")
	}
	songs, err := sc.songservice.GetAlbumSongs(albumId)
	return songs, err
}
