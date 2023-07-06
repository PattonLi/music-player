package controller

import (
	"fmt"
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArtistController struct {
	artistservice *service.ArtistService
}

func NewArtistController() *ArtistController {
	a := service.ArtistService{}
	return &ArtistController{
		artistservice: &a,
	}
}

// 获取十个歌手信息
func (ac *ArtistController) GetTenArtistHandler(c *gin.Context) []model.ArtistInfo {
	artist := ac.artistservice.GetTenAritists()
	return artist
}

// 根据歌手id获取歌手详情
func (ac *ArtistController) GetArtistDetailHandler(c *gin.Context) (model.ArtistInfo, error) {
	id := c.Query("artistId")
	Id, err0 := strconv.Atoi(id)
	if err0 != nil {
		fmt.Println("字符串转换错误")
	}
	artist, err := ac.artistservice.GetArtistDetail(Id)
	return artist, err
}

// 获取歌手描述
func (ac *ArtistController) GetArtistDescribeHandler(c *gin.Context) (string, error) {
	id := c.Query("artistId")
	Id, _ := strconv.Atoi(id)
	profile, err := ac.artistservice.GetArtistDescribe(Id)
	return profile, err
}

// 根据关键词获取歌手
func (ac *ArtistController) GetArtistByKeyWordHandler(c *gin.Context) ([]model.ArtistInfo, error, error, int) {
	keyword := c.Query("keyWord")
	pagesize := c.Query("pageSize")
	currentpage := c.Query("currentPage")

	pageSize, _ := strconv.Atoi(pagesize)
	currntPage, _ := strconv.Atoi(currentpage)

	artistpage, err0, err1, pagetotal := ac.artistservice.GetArtistByKeyWord(pageSize, currntPage, keyword)
	return artistpage, err0, err1, pagetotal
}

// 获取对应筛选条件的歌手
func (ac *ArtistController) GetAtistByCHanddler(c *gin.Context) ([]model.ArtistInfo, error, error, int) {
	pagesize := c.Query("pageSize")
	currentpage := c.Query("currentPage")
	first_letter := c.Query("firstLetter")
	gender := c.Query("gender")
	location := c.Query("location")

	pageSize, _ := strconv.Atoi(pagesize)
	currentPage, _ := strconv.Atoi(currentpage)
	Gender, _ := strconv.Atoi(gender)
	Location, _ := strconv.Atoi(location)

	artistpage, err0, err1, pagetotal := ac.artistservice.GetAtrtistByContition(pageSize, currentPage, first_letter, Gender, Location)
	return artistpage, err0, err1, pagetotal
}

// 获取特定页歌手信息
func (ac *ArtistController) AllArtistInfoHandler(c *gin.Context) ([]model.ArtistInfo, int64) {
	page, _ := strconv.Atoi(c.Query("currentPage"))
	pagesize, _ := strconv.Atoi(c.Query("pageSize"))
	artistlist, totalPage := ac.artistservice.AllArtistInfo(page, pagesize)
	return artistlist, totalPage
}

// 获取特定名称歌手信息
func (ac *ArtistController) ArtistInfoHandler(c *gin.Context) ([]model.ArtistInfo, error) {
	name := c.Query("name")
	artistlist, err := ac.artistservice.ArtistInfo(name)
	return artistlist, err
}

// 添加歌手
func (ac *ArtistController) AddArtistHandler(c *gin.Context) (int64, int64, []model.ArtistInfo, error) {
	artist := model.ArtistInfo{}
	c.BindJSON(&artist)
	totals, currentPage, artistlist, err := ac.artistservice.AddArtistInfo(&artist)
	return totals, currentPage, artistlist, err
}

// 删除歌手信息
func (ac *ArtistController) DeleteArtistInfoHandler(c *gin.Context) error {
	artistID, _ := strconv.Atoi(c.Query("artistId"))
	err := ac.artistservice.DeleteArtistInfo(artistID)
	return err
}

// 歌手信息修改
func (ac *ArtistController) ModifyArtistInfoHandler(c *gin.Context) error {
	artist := model.ArtistInfo{}
	c.BindJSON(&artist)
	err := ac.artistservice.ModifyArtistInfo(&artist)
	return err
}

// 根据专辑id获取歌手
func (ac *ArtistController) GetArtistInfoByalbumIdHanderler(c *gin.Context) (model.ArtistInfo, error) {
	album_id := c.Query("album_id")
	albumIDInt, err := strconv.Atoi(album_id)
	if err != nil {
		fmt.Println("字符串转换错误")
	}
	artist, err := ac.artistservice.GetArtistByAlbumid(albumIDInt)
	return artist, err
}
