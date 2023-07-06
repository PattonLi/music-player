package controller

import (
	"fmt"
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *service.CommentService
	userService    *service.UserService
}

func NewCommentController() *CommentController {
	a := service.CommentService{}
	u := service.UserService{}
	return &CommentController{
		commentService: &a,
		userService:    &u,
	}
}

// 发表评论
func (cc *CommentController) PublishCommentHandler(c *gin.Context) error {
	var err error
	var comment model.CommentInfo
	err = c.ShouldBind(&comment)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	fmt.Println(comment.Comment)
	fmt.Println(comment.User_id)
	fmt.Println(comment.Song_id)
	err = cc.commentService.PublishComment(comment.Comment, comment.User_id, comment.Song_id)
	return err
}

// 获取歌曲所有评论
func (cc *CommentController) GetAllCommentHandler(c *gin.Context) ([]model.Comments, error) {
	songid := c.Query("songId")
	songId, _ := strconv.Atoi(songid)
	comments, err := cc.commentService.GetAllComment(songId)
	return comments, err
}

// 点赞评论
func (cc *CommentController) LikeCommentHandler(c *gin.Context) error {

	var comment model.CommentInfo
	c.ShouldBind(&comment)

	err := cc.commentService.LikeComment(comment.Comment_id, comment.User_id)
	return err
}

// 获取用户所有点赞评论
func (cc *CommentController) GetAllLikeHandler(c *gin.Context) ([]int, error) {
	userid := c.Query("userId")
	userId, _ := strconv.Atoi(userid)

	like_ids, err := cc.commentService.GetAllLike(userId)
	return like_ids, err
}

// 取消点赞评论
func (cc *CommentController) DeleLike(c *gin.Context) error {
	var comment model.CommentInfo
	c.ShouldBind(&comment)

	err := cc.commentService.DeleLike(comment.User_id, comment.Comment_id)
	return err
}
