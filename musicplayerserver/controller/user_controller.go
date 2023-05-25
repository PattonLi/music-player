package controller

import (
	"errors"
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
func (usc *UserController) AddUserHandler(c *gin.Context) error {
	username := c.PostForm("username")
	gender := c.PostForm("gender")
	age := c.PostForm("age")
	email := c.PostForm("email")
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")
	user := model.UserInfo{
		Username: username,
		Gender:   gender,
		Age:      age,
		Email:    email,
		Password: password,
		Nickname: nickname,
		Admin:    "false",
	}
	err := usc.userservice.UserRegister(&user)
	return err
}

// 用户登录
func (usc *UserController) UserLoginHandler(c *gin.Context) (string, uint, error) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := model.UserInfo{
		Username: username,
		Password: password,
	}
	userID, admin, err := usc.userservice.UserLogin(&user)
	if err == nil {
		if admin == "true" { //向用户隐藏管理员账户
			err = errors.New("用户名不存在或密码错误！")
		}
	}
	return username, userID, err
}

// 获取用户信息
func (usc *UserController) UserInfoHandler(c *gin.Context) (*model.UserInfo, error) {
	value, ok := c.Get("admin")
	var user *model.UserInfo
	user = nil
	var err error
	admin, ok2 := value.(string)
	if !ok {
		err = errors.New("获取不到admin验证信息！")
		return user, err
	}
	if !ok2 {
		err = errors.New("admin转化失败！")
		return user, err
	}
	if admin == "ture" {
		err = errors.New("无此用户信息！")
		return user, err
	}
	userID64, _ := strconv.ParseUint(c.Query("userid"), 10, 32)
	userID := uint(userID64)
	user, err = usc.userservice.UserInfo(userID)
	return user, err
}

// 管理员登录
func (usc *UserController) AdminLoginHandler(c *gin.Context) (string, uint, error) {
	adminname := c.PostForm("adminname")
	password := c.PostForm("password")
	adminuser := model.UserInfo{
		Username: adminname,
		Password: password,
	}
	adminID, admin, err := usc.userservice.UserLogin(&adminuser)
	if err == nil {
		if admin == "false" {
			err = errors.New("非管理员账户！")
		}
	}
	return adminname, adminID, err
}

func (usc *UserController) AdminProfileHandler(c *gin.Context) (*model.UserInfo, error) {
	value, ok := c.Get("admin")
	var adminuser *model.UserInfo
	adminuser = nil
	var err error
	admin, ok2 := value.(string)
	if !ok {
		err = errors.New("获取不到admin验证信息！")
		return adminuser, err
	}
	if !ok2 {
		err = errors.New("admin转化失败！")
		return adminuser, err
	}
	if admin == "false" {
		err = errors.New("非管理员请求！")
		return adminuser, err
	}
	fmt.Print(c.Query("adminid"))
	adminID64, _ := strconv.ParseUint(c.Query("adminid"), 10, 32)
	adminID := uint(adminID64)
	adminuser, err = usc.userservice.UserInfo(adminID)
	return adminuser, err
}
