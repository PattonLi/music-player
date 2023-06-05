package controller

import (
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"
	"strconv"
	"github.com/gin-gonic/gin"
)

type AdminUserController struct {
	adminuserservice *service.AdminUserService
}

func NewAdminUserController() *AdminUserController {
	auss := service.NewAdminUserService()
	return &AdminUserController{
		adminuserservice: auss,
	}
}

// 管理员登录
func (ausc *AdminUserController) AdminLoginHandler(c *gin.Context) (string, uint, error) {
	adminname := c.PostForm("adminname")
	password := c.PostForm("password")
	adminuser := model.AdminUserInfo{
		Adminname: adminname,
		Password: password,
	}
	adminID, err := ausc.adminuserservice.AdminUserLogin(&adminuser)
	return adminname, adminID, err
}

//获取特定管理员信息
func (ausc *AdminUserController) AdminInfoHandler(c *gin.Context) (*model.AdminUserInfo, error) {
	adminname := c.Query("adminname")
	adminuser, err := ausc.adminuserservice.AdminUserInfo(adminname)
	return adminuser, err
}

//获取所有管理员信息
func (ausc *AdminUserController) AllAdminInfoHandler() []model.AdminUserInfo{
	adminlist := ausc.adminuserservice.AllAdminInfo()
	return adminlist
}

//管理员信息修改
func (ausc *AdminUserController) ModifyAdminUserInfoHandler(c *gin.Context) error {
	adminuser := model.AdminUserInfo{}
	c.BindJSON(&adminuser)
	err := ausc.adminuserservice.ModifyAdminUserInfo(&adminuser)
	return err
}

//添加管理员信息
func (ausc *AdminUserController) AddAdminUserHandler(c *gin.Context) (error) {
	adminuser := model.AdminUserInfo{}
	c.BindJSON(&adminuser)
	err := ausc.adminuserservice.AddAdminUserInfo(&adminuser)
	return err
}

//删除管理员信息
func (ausc *AdminUserController) DeleteAdminUserInfoHandler(c *gin.Context) error {
	adminID,_ := strconv.Atoi(c.Query("adminId"))
	err := ausc.adminuserservice.DeleteAdminUserInfo(adminID)
	return err
}