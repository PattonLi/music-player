package controller

import (
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlaylistController struct {
	playlistservice *service.PlaylistService
}

func NewPlaylistController() *PlaylistController {
	p := service.PlaylistService{}
	return &PlaylistController{
		playlistservice: &p,
	}
}

// 获取热门歌单
func (pc *PlaylistController) GetHotPlaylistHandler(c *gin.Context) ([]model.PlaylistInfo, error, error, int) {
	pagesize := c.Query("pageSize")
	currentpage := c.Query("currentPage")
	pageSize, _ := strconv.Atoi(pagesize)
	currentPage, _ := strconv.Atoi(currentpage)

	playlistpage, err0, err1, pagetotal := pc.playlistservice.GetHotPlaylist(pageSize, currentPage)
	return playlistpage, err0, err1, pagetotal
}

// 获取歌单所有歌曲
func (pc *PlaylistController) GetPlaylistAllSongsHandler(c *gin.Context) (model.PlaylistInfo, []model.SongInfo, bool) {
	playlist_id := c.Query("playListId")
	playList_id, _ := strconv.Atoi(playlist_id)
	playlist, songs, err := pc.playlistservice.GetPlaylistAllSongs(playList_id)
	return playlist, songs, err
}
