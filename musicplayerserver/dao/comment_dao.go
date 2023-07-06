package dao

import (
	"errors"
	"music-player/musicplayerserver/model"
	"strconv"

	"gorm.io/gorm"
)

type CommentDao struct {
}

// 发表评论,点赞评论
func (cd *CommentDao) InsertComment(comment model.CommentInfo) error {
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

//获取特定歌曲评论
func (cd *CommentDao) GetSongCommentById(songID int)([]model.CommentInfo, error) {
	var commentlist []model.CommentInfo
	err := DB.Where("song_id = ? ", songID).Where("type = ?", 1).Find(&commentlist).Error
	if err != nil && !errors.Is(err,gorm.ErrRecordNotFound){
		err = errors.New("数据库查找错误！")
	}
	return commentlist, err
}

//删除特定评论
func (cd *CommentDao) DeleteCommentById(commentID int) error {
	err := DB.Delete(&model.CommentInfo{}, commentID).Error
	return err
}
