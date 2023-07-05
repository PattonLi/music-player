package service

import (
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
	"time"
)

type LogService struct {
	logdao *dao.LogDao
}

func (l *LogService) AddPlayLog(log model.LogInfo) error {
	err := l.logdao.AddPlayLog(log)
	return err
}

func (l *LogService) AddRegisterLog(userID int, username string) error {
	var log model.LogInfo
	log.User_id = userID
	log.Username = username
	log.Time = time.Now().Format("2006-01-02 15:04:05")
	log.Type = 1
	err := l.logdao.AddPlayLog(log)
	return err
}
