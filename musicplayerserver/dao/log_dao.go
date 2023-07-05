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

//根据type获得日志
func (l *LogDao) GetLogByType(logtype int) ([]model.LogInfo,error) {
	var loglist []model.LogInfo
	err := DB.Where("type = ?", logtype).Find(&loglist).Error
	return loglist, err
}
