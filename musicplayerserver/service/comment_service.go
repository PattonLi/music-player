package service

import (
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
	"time"
)

type CommentService struct {
	commentDao *dao.CommentDao
}

func NewCommentService() *CommentService {
	c := dao.CommentDao{}
	return &CommentService{
		commentDao: &c,
	}
}

// 发表评论
func (cs *CommentService) PublishComment(comment string, userId int, songId int) error {
	err := cs.commentDao.InsertComment(model.CommentInfo{
		User_id:      userId,
		Song_id:      songId,
		Comment:      comment,
		Like:         0,
		Comment_time: time.RFC1123,
		Type:         1,
	})
	return err
}
