package dao

import "music-player/musicplayerserver/model"

type CommentDao struct {
}

// 发表评论
func (cd *CommentDao) InsertComment(comment model.CommentInfo) error {
	result := DB.Create(comment)
	println(result.Error)
	println("exec dao")
	return result.Error
}
