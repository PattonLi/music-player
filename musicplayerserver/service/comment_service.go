package service

import (
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
	"strconv"
	"time"
)

type CommentService struct {
	commentDao *dao.CommentDao
	userDao    *dao.UserDao
}

func NewCommentService() *CommentService {
	c := dao.CommentDao{}
	u := dao.UserDao{}
	return &CommentService{
		commentDao: &c,
		userDao:    &u,
	}
}

// 发表评论
func (cs *CommentService) PublishComment(comment string, userId int, songId int) error {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	err := cs.commentDao.InsertComment(model.CommentInfo{
		User_id:      userId,
		Song_id:      songId,
		Comment:      comment,
		Like:         0,
		Comment_time: formattedTime,
		Type:         1,
	})
	return err
}

// 获取歌曲所有评论
func (cs *CommentService) GetAllComment(songid int) ([]model.Comments, error) {
	comment, err := cs.commentDao.GetAllComment(songid)
	println(len(comment))
	var comments []model.Comments
	for _, c := range comment {
		var cm model.Comments
		var user *model.UserInfo
		var nickname string
		var picurl string
		userid := c.User_id
		user, _ = cs.userDao.GetUserProfile(userid)
		nickname = user.Nickname
		picurl = user.Pic_url

		cm.Nickname = nickname
		cm.Picurl = picurl
		cm.Comment = c.Comment
		cm.Commentid = c.Comment_id
		cm.Commenttime = c.Comment_time
		cm.Like = c.Like

		comments = append(comments, cm)
	}

	return comments, err
}

// 点赞评论
func (cs *CommentService) LikeComment(commentId int, userId int) error {

	err := cs.commentDao.LikeComment(model.CommentInfo{
		User_id:      userId,
		Type:         2,
		Comment_time: time.RFC1123,
		Comment:      strconv.Itoa(commentId),
	})
	return err
}

// 获得用户所有点赞评论
func (cs *CommentService) GetAllLike(userid int) ([]int, error) {
	like_ids, err := cs.commentDao.GetAllLike(userid)
	return like_ids, err
}

// 取消点赞评论
func (cs *CommentService) DeleLike(userid int, comment_id int) error {
	err := cs.commentDao.DeleLike(userid, comment_id)
	return err
}

// 获取特定歌曲评论
func (cs *CommentService) GetSongComment(songID int) ([]model.CommentInfo, error) {
	commentlist, err := cs.commentDao.GetSongCommentById(songID)
	return commentlist, err
}

// 删除特定评论
func (cs *CommentService) DeleteComment(commentID int) error {
	err := cs.commentDao.DeleteCommentById(commentID)
	return err
}
