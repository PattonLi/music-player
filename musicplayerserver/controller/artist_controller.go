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
