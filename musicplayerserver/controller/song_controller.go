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

// 分页获取歌手歌曲
func (sc *SongController) GetSongsPageHandler(c *gin.Context) ([]model.SongInfo, error, error, int) {
	params := []string{"artistId", "order", "currentPage", "pageSize"}
	id := c.Query(params[0])
	order := c.Query(params[1])
	currentpage := c.Query(params[2])
	pagesize := c.Query(params[3])

	Id, _ := strconv.Atoi(id)
	currentPage, _ := strconv.Atoi(currentpage)
	pageSize, _ := strconv.Atoi(pagesize)

	songpage, err0, err1, pagetotal := sc.songservice.GetSongsPage(Id, order, currentPage, pageSize)
	return songpage, err0, err1, pagetotal
}

// 根据songid获取歌曲
func (sc *SongController) GetSongDetailHandler(c *gin.Context) (model.SongInfo, error) {
	id := c.Query("songId")
	Id, _ := strconv.Atoi(id)
	song, err := sc.songservice.GetSongDetail(Id)
	return song, err
}

// 根据关键词获取歌曲
func (sc *SongController) GetSongByKeyWordHandler(c *gin.Context) ([]model.SongInfo, error, error, int) {
	keyword := c.Query("keyWord")
	pagesize := c.Query("pageSize")
	currentpage := c.Query("currentPage")

	pageSize, _ := strconv.Atoi(pagesize)
	currntPage, _ := strconv.Atoi(currentpage)

	songpage, err0, err1, pagetotal := sc.songservice.GetSongByKeyWord(pageSize, currntPage, keyword)
	return songpage, err0, err1, pagetotal
}
