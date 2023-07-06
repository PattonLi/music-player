package dao

import (
	"errors"
	"fmt"
	"music-player/musicplayerserver/model"

	"gorm.io/gorm"
)

type CommentDao struct {
}

// 发表评论
func (cd *CommentDao) InsertComment(comment model.CommentInfo) error {
	result := DB.Create(comment)
	println(result.Error)
	println("exec dao")
	return result.Error
}

//获取特定歌曲评论
func (cd *CommentDao) GetCommentBySongID(songID int) ([]model.CommentInfo, error) {
	var commentlist []model.CommentInfo
	err := DB.Where("song_id = ?", songID).Where("type = ?", 1).Find(&commentlist).Error
	if (err != nil && !errors.Is(err, gorm.ErrRecordNotFound)){
		err = errors.New("数据库查找出错！")
		fmt.Println(err.Error())
	}
	return commentlist, err
}

//删除特定评论
func (cd *CommentDao) DeleteCommentByID(commentID int) error {
	err := DB.Delete(&model.CommentInfo{}, commentID).Error
	return err
}
