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
