package controller

import (
	"fmt"
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	albumService *service.AlbumService
}

func NewAlbumController() *AlbumController {
	al := service.AlbumService{}
	return &AlbumController{
		albumService: &al,
	}
}

// 获取十张专辑
func (ac *AlbumController) GetTenAlbumsHandler(c *gin.Context) []model.AlbumInfo {
	albums := ac.albumService.GetTenAlbums()
	return albums
}

// 根据专辑id获取专辑
func (ac *AlbumController) GetAlbumByIdHandler(c *gin.Context) (model.AlbumInfo, error) {
	id := c.Query("albumId")
	Id, err0 := strconv.Atoi(id)
	if err0 != nil {
		fmt.Println("字符串转换错误")
	}

	album, err := ac.albumService.GetAlbumById(Id)
	return album, err
}

// 获取特定名称专辑信息
func (ac *AlbumController) AlbumInfoHandler(c *gin.Context) ([]gin.H, error) {
	name := c.Query("name")
	albumlist, err := ac.albumService.AlbumInfo(name)
	albums := make([]gin.H, 0)
	for _, album := range albumlist {
		albuminfo := gin.H{
			"albumId":  album.AlbumID,
			"size":     album.Size,
			"name":     album.Name,
			"artist":   album.Artist,
			"ArtistID": album.Artist_ID,
			"profile":  album.Profile,
		}
		albums = append(albums, albuminfo)
	}
	return albums, err
}

// 获取特定页所有专辑信息
func (ac *AlbumController) AllAlbumInfoHandler(c *gin.Context) ([]model.AlbumInfo, int64) {
	page, _ := strconv.Atoi(c.Query("currentPage"))
	pagesize, _ := strconv.Atoi(c.Query("pageSize"))
	albumlist, totalPage := ac.albumService.AllAlbumInfo(page, pagesize)
	return albumlist, totalPage
}
