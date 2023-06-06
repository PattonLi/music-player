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
func (ac *AlbumController) AlbumInfoHandler(c *gin.Context) ([]model.AlbumInfo, error) {
	name := c.Query("name")
	albumlist, err := ac.albumService.AlbumInfo(name)
	return albumlist, err
}

// 获取特定页所有专辑信息
func (ac *AlbumController) AllAlbumInfoHandler(c *gin.Context) ([]model.AlbumInfo, int64) {
	page, _ := strconv.Atoi(c.Query("currentPage"))
	pagesize, _ := strconv.Atoi(c.Query("pageSize"))
	albumlist, totalPage := ac.albumService.AllAlbumInfo(page, pagesize)
	return albumlist, totalPage
// 分页获取歌手专辑
func (ac *AlbumController) GetAlbumPageHandler(c *gin.Context) ([]model.AlbumInfo, error, error, int) {
	artistid := c.Query("artistId")
	currentpage := c.Query("currentPage")
	pagesize := c.Query("pageSize")

	artistId, _ := strconv.Atoi(artistid)
	currentPage, _ := strconv.Atoi(currentpage)
	pageSize, _ := strconv.Atoi(pagesize)

	albumpage, err0, err1, pagenum := ac.albumService.GetAlbumPage(artistId, currentPage, pageSize)
	return albumpage, err0, err1, pagenum
}
