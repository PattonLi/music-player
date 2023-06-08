package controller

import (
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LikesController struct {
	likeservice *service.LikesService
}

func NewLikesController() *LikesController {
	l := service.LikesService{}
	return &LikesController{
		likeservice: &l,
	}
}

func (lc *LikesController) GetUserLikesHandler(c *gin.Context) ([]model.SongInfo, []model.AlbumInfo, []model.ArtistInfo, []model.PlaylistInfo, error) {
	userid := c.Query("userId")
	userId, _ := strconv.Atoi(userid)
	songs, albums, artists, playlists, err := lc.likeservice.GetUserLikes(userId)
	return songs, albums, artists, playlists, err
}
