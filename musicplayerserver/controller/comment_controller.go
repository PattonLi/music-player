package controller

import (
	"fmt"
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentController() *CommentController {
	a := service.CommentService{}
	return &CommentController{
		commentService: &a,
	}
}

// 发表评论
func (cc *CommentController) PublishCommentHandler(c *gin.Context) error {
	println("exec comment conteoller")
	var err error
	var comment model.CommentInfo
	err = c.ShouldBind(&comment)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	err = cc.commentService.PublishComment(comment.Comment, comment.User_id, comment.Song_id)
	return err
}
