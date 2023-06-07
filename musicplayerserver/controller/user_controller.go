package controller

import (
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"
	"strconv"
	"music-player/musicplayerserver/utils/jwt"
	"github.com/gin-gonic/gin"
	"errors"
	"strings"
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

// 添加用户
func (usc *UserController) AddUserHandler(c *gin.Context) (*model.UserInfo, error) {
	user := model.UserInfo{}
	c.BindJSON(&user)
	err := usc.userservice.AddUserInfo(&user)
	return &user, err
}

//用户信息修改
func (usc *UserController) ModifyUserInfoHandler(c *gin.Context) error {
	user := model.UserInfo{}
	c.BindJSON(&user)
	err := usc.userservice.ModifyUserInfo(&user)
	return err
}

//删除用户信息
func (usc *UserController) DeleteUserInfoHandler(c *gin.Context) error {
	userID,_ := strconv.Atoi(c.Query("userId"))
	err := usc.userservice.DeleteUserInfo(userID)
	return err
}

// 获取特定名称用户信息
func (usc *UserController) UserInfoHandler(c *gin.Context) ([]model.UserInfo, error) {
	name := c.Query("name")
	userlist, err := usc.userservice.UserInfo(name)
	return userlist, err
}

// 获取特定页所有用户信息
func (usc *UserController) AllUserInfoHandler(c *gin.Context) ([]model.UserInfo, int64) {
	page, _ := strconv.Atoi(c.Query("currentPage"))
	pagesize, _ := strconv.Atoi(c.Query("pageSize"))
	userlist, totalrecord := usc.userservice.AllUserInfo(page, pagesize)
	return userlist, totalrecord
}

//根据ID获取单个用户信息
func (usc *UserController) UserProfileHandler(c *gin.Context) (*gin.H, error) {
	userID,_ := strconv.Atoi(c.Query("userId"))
	user, err := usc.userservice.UserProfile(userID)
	userprofile := gin.H{
		"nickname": user.Nickname,
        "username": user.Username,
        "picUrl": user.Pic_url,
        "phone": user.Phone,
        "email": user.Email,
        "age": user.Age,
        "gender": user.Gender,
	}
	return &userprofile,err
}

// 用户手机号登录
func (usc *UserController) UserLoginHandler(c *gin.Context) (int, string, error) {
	phone := c.Query("phone")
	password := c.Query("password")
	user := model.UserInfo{
		Phone: phone,
		Password: password,
	}
	userID, err := usc.userservice.UserLogin(&user)
	var token string = ""
	if err == nil {
		token,_ = utils.CreateToken(userID)
	}
	return userID, token, err
}

//用户手机号注册
func (usc *UserController) UserRegisterHandler(c *gin.Context) (int, string, error) {
	user := model.UserInfo{}
	c.BindJSON(&user)
	userID, err := usc.userservice.UserRegister(&user)
	var token string = ""
	if err == nil {
		token,_ = utils.CreateToken(userID)
	}
	return userID, token, err
}

//上传用户头像
func (usc *UserController) UploadUserPicHandler(c *gin.Context) error{
	fileHeader, err := c.FormFile("file")
    if err != nil {
		err = errors.New("照片上传失败！")
        return err
    }
	if !strings.HasPrefix(fileHeader.Header.Get("Content-Type"), "image/") {
        err = errors.New("上传文件非图片！")
        return err
    }
	temp,_ := c.Get("id")
	userID,_ := temp.(int)
	err = usc.userservice.UploadUserPic(userID, fileHeader)
	return err
}
