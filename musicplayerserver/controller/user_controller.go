package controller

import (
	"errors"
	"errors"
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"
	utils "music-player/musicplayerserver/utils/jwt"
	utils "music-player/musicplayerserver/utils/jwt"
	"strconv"
	"strings"

	"strings"

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

// 添加用户
func (usc *UserController) AddUserHandler(c *gin.Context) (int64, int64, []model.UserInfo, error) {
	user := model.UserInfo{}
	c.BindJSON(&user)
	totals, currentPage, userlist, err := usc.userservice.AddUserInfo(&user)
	return totals, currentPage, userlist, err
}

// 用户信息修改
func (usc *UserController) ModifyUserInfoHandler(c *gin.Context) error {
	user := model.UserInfo{}
	c.BindJSON(&user)
	err := usc.userservice.ModifyUserInfo(&user)
	return err
}

// 删除用户信息
// 删除用户信息
func (usc *UserController) DeleteUserInfoHandler(c *gin.Context) error {
	userID, _ := strconv.Atoi(c.Query("user_id"))
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

// 根据ID获取单个用户信息
// 根据ID获取单个用户信息
func (usc *UserController) UserProfileHandler(c *gin.Context) (*gin.H, error) {
	userID, _ := strconv.Atoi(c.Query("userId"))
	userID, _ := strconv.Atoi(c.Query("userId"))
	user, err := usc.userservice.UserProfile(userID)
	userprofile := gin.H{
		"nickname": user.Nickname,
		"username": user.Username,
		"picUrl":   user.Pic_url,
		"phone":    user.Phone,
		"email":    user.Email,
		"age":      strconv.Itoa(user.Age),
		"gender":   user.Gender,
	}
	return &userprofile, err
}

//g根据index获取单个用户信息
func (usc *UserController) GetAInfoHandler(c *gin.Context) (int64, *model.UserInfo, error) {
	index,_ := strconv.Atoi(c.Query("index"))
	totals, user, err := usc.userservice.AUserInfo(index)
	return totals, user, err
}

// 用户手机号登录
func (usc *UserController) UserLoginHandler(c *gin.Context) (int, string, error) {
	phone := c.Query("phone")
	password := c.Query("password")
	user := model.UserInfo{
		Phone:    phone,
		Phone:    phone,
		Password: password,
	}
	userID, err := usc.userservice.UserLogin(&user)
	var token string = ""
	if err == nil {
		token, _ = utils.CreateToken(userID)
		token, _ = utils.CreateToken(userID)
	}
	return userID, token, err
}

// 用户手机号注册
func (usc *UserController) UserRegisterHandler(c *gin.Context) (int, string, error) {
	user := model.UserInfo{}
	c.BindJSON(&user)
	userID, err := usc.userservice.UserRegister(&user)
	var token string = ""
	if err == nil {
		token, _ = utils.CreateToken(userID)
	}
	return userID, token, err
}*/

// 用户自行上传头像
func (usc *UserController) UploadUserPicHandler(c *gin.Context) (string, error) {
	fileHeader, err := c.FormFile("file")
	var url string = ""
	if err != nil {
		err = errors.New("照片上传失败！")
		return url, err
	}
	if !strings.HasPrefix(fileHeader.Header.Get("Content-Type"), "image/") {
		err = errors.New("上传文件非图片！")
		return url, err
	}
	temp, _ := c.Get("id")
	userID, _ := temp.(int)
	url, err = usc.userservice.UploadUserPic(userID, fileHeader)
	return url, err
}

//管理员上传用户头像
func (usc *UserController) AdminUploadUserPicHandler(c *gin.Context) (string, error) {
	fileHeader, err := c.FormFile("file")
	var url string = ""
	if err != nil {
		err = errors.New("照片上传失败！")
		return url, err
	}
	if !strings.HasPrefix(fileHeader.Header.Get("Content-Type"), "image/") {
		err = errors.New("上传文件非图片！")
		return url, err
	}
	userID,_ := strconv.Atoi(c.PostForm("userId"))
	url, err = usc.userservice.UploadUserPic(userID, fileHeader)
	return url, err
}
