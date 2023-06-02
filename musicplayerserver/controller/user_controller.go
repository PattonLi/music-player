package controller

import (
	"errors"
	"fmt"
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"

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

// 获取单个用户信息
func (usc *UserController) UserInfoHandler(c *gin.Context) (*model.UserInfo, error) {
	var user *model.UserInfo
	user = nil
	var err error
	/*value, ok := c.Get("admin")
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
	}*/
	username := c.Query("username")
	user, err = usc.userservice.UserInfo(username)
	return user, err
}

//获取所有用户信息
func (usc *UserController) AllUserInfoHandler() []gin.H{
	userlist := usc.userservice.AllUserInfo()
	users := make([]gin.H,0)
	for _, user := range userlist{
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
		}
		users = append(users,userinfo)
	}
	return users
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

//获取特定管理员信息
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
	adminname := c.Query("adminname")
	adminuser, err = usc.userservice.UserInfo(adminname)
	return adminuser, err
}

//获取所有管理员信息
func (usc *UserController) AllAdminInfoHandler() []gin.H{
	adminlist := usc.userservice.AllAdminInfo()
	adminusers := make([]gin.H,0)
	for _, adminuser := range adminlist{
		userinfo := gin.H{
			"user_id":    adminuser.ID,
			"username":   adminuser.Username,
			"password":   adminuser.Password,
		}
		adminusers = append(adminusers,userinfo)
	}
	return adminusers
}
