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
	name := c.Query("albumName")
	albumlist, err := ac.albumService.AlbumInfo(name)
	return albumlist, err
}

// 获取特定页所有专辑信息
func (ac *AlbumController) AllAlbumInfoHandler(c *gin.Context) ([]model.AlbumInfo, int64) {
	page, _ := strconv.Atoi(c.Query("currentPage"))
	pagesize, _ := strconv.Atoi(c.Query("pageSize"))
	albumlist, totalrecord := ac.albumService.AllAlbumInfo(page, pagesize)
	return albumlist, totalrecord
}

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

// 根据关键词获取专辑
func (ac *AlbumController) GetAlbumByKeyWordHandler(c *gin.Context) ([]model.AlbumInfo, error, error, int) {
	keyword := c.Query("keyWord")
	pagesize := c.Query("pageSize")
	currentpage := c.Query("currentPage")

	pageSize, _ := strconv.Atoi(pagesize)
	currntPage, _ := strconv.Atoi(currentpage)

	albumpage, err0, err1, pagetotal := ac.albumService.GetAlbumByKeyWord(pageSize, currntPage, keyword)
	return albumpage, err0, err1, pagetotal
}

// 添加专辑
func (ac *AlbumController) AddAlbumHandler(c *gin.Context) (int64, int64, []model.AlbumInfo, error) {
	album := model.AlbumInfo{}
	c.BindJSON(&album)
	totals, currentPage, albumlist, err := ac.albumService.AddAlbumInfo(&album)
	return totals, currentPage, albumlist, err
}

// 专辑信息修改
func (alc *AlbumController) ModifyAlbumInfoHandler(c *gin.Context) error {
	album := model.AlbumInfo{}
	c.BindJSON(&album)
	err := alc.albumService.ModifyAlbumInfo(&album)
	return err
}

// 删除专辑信息
func (alc *AlbumController) DeleteAlbumInfoHandler(c *gin.Context) error {
	albumID, _ := strconv.Atoi(c.Query("albumId"))
	err := alc.albumService.DeleteAlbumInfo(albumID)
	return err
}

// 根据歌手id获取专辑
func (alc *AlbumController) GetAlbumInfoByartistIdHanderler(c *gin.Context) ([]model.AlbumInfo, error) {
	artist_id := c.Query("artistId")
	artistIDInt, err := strconv.Atoi(artist_id)
	if err != nil {
		fmt.Println("字符串转换错误")
	}
	albumlist, err := alc.albumService.GetAlbumByArtistid(artistIDInt)
	return albumlist, err
}
