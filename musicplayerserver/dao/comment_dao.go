package dao

import (
	"music-player/musicplayerserver/model"
	"strconv"
)

type CommentDao struct {
}

// 发表评论,点赞评论
func (cd *CommentDao) InsertComment(comment model.CommentInfo) error {
	var song model.SongInfo
	songid := comment.Song_id
	DB.First(&song, songid)
	mark := song.Mark
	DB.Model(&song).Update("mark", mark+1)

	result := DB.Create(&comment)
	return result.Error
}

// 获取歌曲所有评论
func (cd *CommentDao) GetAllComment(songid int) ([]model.CommentInfo, error) {
	var comment []model.CommentInfo
	result := DB.Where("song_id = ?", songid).Find(&comment)
	return comment, result.Error
}

// 点赞评论
func (cd *CommentDao) LikeComment(comment model.CommentInfo) error {
	var c model.CommentInfo

	result1 := DB.Create(&comment)
	comment_id, _ := strconv.Atoi(comment.Comment)
	result2 := DB.Where("comment_id = ? AND type = ?", comment_id, 1).Find(&c)

	likes := c.Like
	result3 := DB.Model(&model.CommentInfo{}).Where("comment_id = ?", comment_id).Update("like", likes+1)

	if result1.Error != nil {
		return result1.Error
	}

	if result2.Error != nil {
		return result2.Error
	}

	if result3.Error != nil {
		return result3.Error
	}

	return nil
}

// 获得用户所有点赞评论
func (cd *CommentDao) GetAllLike(userid int) ([]int, error) {
	var comment []model.CommentInfo
	result := DB.Where("user_id = ? AND type = ?", userid, 2).Find(&comment)

	var like_ids []int
	for _, c := range comment {
		id, _ := strconv.Atoi(c.Comment)
		like_ids = append(like_ids, id)
	}
	return like_ids, result.Error
}

// 取消点赞评论
func (cd *CommentDao) DeleLike(userid int, comment_id int) error {
	var c model.CommentInfo
	result2 := DB.Where("comment_id = ? AND type = ?", comment_id, 1).Find(&c)

	likes := c.Like
	result3 := DB.Model(&model.CommentInfo{}).Where("comment_id = ?", comment_id).Update("like", likes-1)

	result1 := DB.Where("comment = ? AND user_id = ?", comment_id, userid).Delete(&model.CommentInfo{})

	if result1.Error != nil {
		return result1.Error
	}

	if result2.Error != nil {
		return result2.Error
	}

	if result3.Error != nil {
		return result3.Error
	}

	return nil
}
