package service

import (
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
)

type LogService struct {
	logdao *dao.LogDao
}

func (l *LogService) AddPlayLog(log model.LogInfo) error {
	err := l.logdao.AddPlayLog(log)
	return err
}
