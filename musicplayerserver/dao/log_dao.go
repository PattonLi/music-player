package dao

import (
	"music-player/musicplayerserver/model"
)

type LogDao struct {
}

func (l *LogDao) AddPlayLog(log model.LogInfo) error {
	result := DB.Create(&log)
	return result.Error
}
