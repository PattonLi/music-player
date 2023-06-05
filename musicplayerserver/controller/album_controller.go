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
