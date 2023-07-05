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
	userService *service.UserService
}

func NewCommentController() *CommentController {
	a := service.CommentService{}
	u := service.UserService{}
	return &CommentController{
		commentService: &a,
		userService: &u,
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
	fmt.Println(comment.Comment)
	fmt.Println(comment.User_id)
	fmt.Println(comment.Song_id)
	err = cc.commentService.PublishComment(comment.Comment, comment.User_id, comment.Song_id)
	return err
}


//获取特定歌曲评论
func (cc *CommentController) GetSongCommentHandler(c *gin.Context) ([]gin.H, error) {
	songID,_ := strconv.Atoi(c.Query("songId"))
	var frontendcommentlist []gin.H
	commentlist, err := cc.commentService.GetSongComment(songID)
	if err != nil {
		return frontendcommentlist,err
	} else {
		err = nil
		for _, comment := range commentlist {
			err = nil
			user,err := cc.userService.UserProfile(comment.User_id)
			if err == nil {
				frontendcomment := gin.H{
					"nickname": user.Nickname,
					"picUrl": user.Pic_url,
					"comment": comment.Comment,
					"commentId": comment.Comment_id,
					"commentTime": comment.Comment_time,
					"like": strconv.Itoa(comment.Like),
				}
				frontendcommentlist = append(frontendcommentlist, frontendcomment)
			}
		}
		err = nil
		return frontendcommentlist, err
	}
}

//删除特定评论
func (cc *CommentController) DeltetCommentHandler(c *gin.Context) error {
	commentID,_ := strconv.Atoi(c.Query("commentId"))
	err := cc.commentService.DeleteComment(commentID)
	return err
}