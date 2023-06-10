package controller

import (
	"errors"
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"
	"time"

	"github.com/gin-gonic/gin"
)

type LogController struct {
	logservice *service.LogService
}

func NewLogController() *LogController {
	l := service.LogService{}
	return &LogController{
		logservice: &l,
	}
}

type Log struct {
	SongId int `json:"songId"`
	UserId int `json:"userId"`
}

func (lc *LogController) AddPlayLogHandler(c *gin.Context) error {

	var log Log
	var playlog model.LogInfo
	songservice := service.SongService{}
	userservice := service.UserService{}
	c.ShouldBind(&log)

	songid := log.SongId
	userid := log.UserId

	song, err0 := songservice.GetSongDetail(songid)
	user, err1 := userservice.UserProfile(userid)

	playlog.Song_id = songid
	playlog.User_id = userid
	playlog.Songname = song.Name
	playlog.Username = user.Username
	playlog.Time = time.Now().Format("2006-01-02 15:04:05")
	playlog.Type = 2

	lc.logservice.AddPlayLog(playlog)

	if err0 == nil && err1 == nil {
		return nil
	} else {
		return errors.New("数据库操作有误")
	}
}

func (lc *LogController) AddDownloadLog(c *gin.Context) error {
	var log Log
	var playlog model.LogInfo
	songservice := service.SongService{}
	userservice := service.UserService{}
	c.ShouldBind(&log)

	songid := log.SongId
	userid := log.UserId

	song, err0 := songservice.GetSongDetail(songid)
	user, err1 := userservice.UserProfile(userid)

	playlog.Song_id = songid
	playlog.User_id = userid
	playlog.Songname = song.Name
	playlog.Username = user.Username
	playlog.Time = time.Now().Format("2006-01-02 15:04:05")
	playlog.Type = 3

	lc.logservice.AddPlayLog(playlog)

	if err0 == nil && err1 == nil {
		return nil
	} else {
		return errors.New("数据库操作有误")
	}
}
