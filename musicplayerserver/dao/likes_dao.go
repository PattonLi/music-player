package dao

import (
	"music-player/musicplayerserver/model"
)

type LikeDao struct {
}

// 获取用户所有收藏
func (l *LikeDao) GetUserLikes(userid int) ([]model.LikesInfo, error) {
	var likes []model.LikesInfo
	result := DB.Where("user_id = ?", userid).Find(&likes)
	return likes, result.Error
}

// 添加用户收藏
func (l *LikeDao) AddUserLikes(like model.LikesInfo) error {
	result := DB.Create(&like)
	return result.Error
}

// 删除用户收藏
func (l *LikeDao) DeleteUserLikes(like model.LikesInfo) error {
	if like.Type == 1 {
		result := DB.Where("user_id = ? AND song_id = ? ", like.User_id, like.Song_id).Delete(&like)
		return result.Error
	}

	if like.Type == 2 {
		result := DB.Where("user_id = ? AND artist_id = ? ", like.User_id, like.Artist_id).Delete(&like)
		return result.Error
	}

	if like.Type == 3 {
		result := DB.Where("user_id = ? AND album_id = ? ", like.User_id, like.Album_id).Delete(&like)
		return result.Error
	}

	result := DB.Where("user_id = ? AND playlist_id = ? ", like.User_id, like.Playlist_id).Delete(&like)
	return result.Error

}
