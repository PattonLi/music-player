package controller

import (
	"fmt"
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userservice *service.UserService
}

func NewUserController() *UserController {
	uss := service.NewUserService()
	return &UserController{
		userservice: uss,
	}
}

// 用户注册
func (usc *UserController) AddUserHandler(c *gin.Context) (*model.UserInfo, error) {
	user := model.UserInfo{}
	c.BindJSON(&user)
	err := usc.userservice.UserRegister(&user)
	return &user, err
}

// 获取特定名称用户信息
func (usc *UserController) UserInfoHandler(c *gin.Context) ([]gin.H, error) {
	name := c.Query("name")
	userlist, err := usc.userservice.UserInfo(name)
	users := make([]gin.H, 0)
	for _, user := range userlist {
		userinfo := gin.H{
			"user_id":    user.ID,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
			"deleted_at": user.DeletedAt,
			"username":   user.Username,
			"gender":     user.Gender,
			"age":        user.Age,
			"email":      user.Email,
			"password":   user.Password,
			"nickname":   user.Nickname,
			"phone":      user.Phone,
			"pic_url":    user.Pic_url,
		}
		users = append(users, userinfo)
	}
	return users, err
}

// 获取特定页所有用户信息
func (usc *UserController) AllUserInfoHandler(c *gin.Context) ([]gin.H, int64) {
	page, _ := strconv.Atoi(c.Query("currentPage"))
	pagesize, _ := strconv.Atoi(c.Query("pageSize"))
	userlist, totalPage := usc.userservice.AllUserInfo(page, pagesize)
	users := make([]gin.H, 0)
	for _, user := range userlist {
		fmt.Println(user.ID)
		userinfo := gin.H{
			"user_id":    user.ID,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
			"deleted_at": user.DeletedAt,
			"username":   user.Username,
			"gender":     user.Gender,
			"age":        user.Age,
			"email":      user.Email,
			"password":   user.Password,
			"nickname":   user.Nickname,
			"phone":      user.Phone,
			"pic_url":    user.Pic_url,
		}
		users = append(users, userinfo)
	}
	return users, totalPage
}

// 用户登录
func (usc *UserController) UserLoginHandler(c *gin.Context) (string, uint, error) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := model.UserInfo{
		Username: username,
		Password: password,
	}
	userID, err := usc.userservice.UserLogin(&user)
	return username, userID, err
}
