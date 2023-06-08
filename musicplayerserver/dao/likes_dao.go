package dao

import (
	"music-player/musicplayerserver/model"
)

type LikeDao struct {
}

func (l *LikeDao) GetUserLikes(userid int) ([]model.LikesInfo, error) {
	var likes []model.LikesInfo
	result := DB.Where("user_id = ?", userid).Find(&likes)
	return likes, result.Error
}
