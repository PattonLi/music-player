package controller

import (
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"
	utils "music-player/musicplayerserver/utils/jwt"
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
func (ausc *AdminUserController) AdminLoginHandler(c *gin.Context) (int, string, error) {
	adminname := c.Query("adminName")
	password := c.Query("password")
	adminuser := model.AdminUserInfo{
		Adminname: adminname,
		Password: password,
	}
	adminID, err := ausc.adminuserservice.AdminUserLogin(&adminuser)
	var token string = ""
	if err == nil {
		token, _ = utils.CreateToken(adminID)
	}
	return adminID, token, err
}

//根据ID获取单个管理员信息
func (ausc *AdminUserController) AdminProfileHandler(c *gin.Context) (string, error) {
	adminID,_ := strconv.Atoi(c.Query("adminId"))
	adminname,err := ausc.adminuserservice.AdminProfile(adminID)
	return adminname, err
}

//获取特定管理员信息
func (ausc *AdminUserController) AdminInfoHandler(c *gin.Context) ([]model.AdminUserInfo, error) {
	adminname := c.Query("adminname")
	adminlist, err := ausc.adminuserservice.AdminUserInfo(adminname)
	return adminlist, err
}

//获取特定页所有管理员信息
func (ausc *AdminUserController) AllAdminInfoHandler(c *gin.Context) (int64, []model.AdminUserInfo){
	page, _ := strconv.Atoi(c.Query("currentPage"))
	pagesize, _ := strconv.Atoi(c.Query("pageSize"))
	totals, adminlist := ausc.adminuserservice.AllAdminInfo(page,pagesize)
	return totals, adminlist
}

//管理员信息修改
func (ausc *AdminUserController) ModifyAdminUserInfoHandler(c *gin.Context) error {
	adminuser := model.AdminUserInfo{}
	c.BindJSON(&adminuser)
	err := ausc.adminuserservice.ModifyAdminUserInfo(&adminuser)
	return err
}

//添加管理员信息
func (ausc *AdminUserController) AddAdminUserHandler(c *gin.Context) (int64, int64, []model.AdminUserInfo,error) {
	adminuser := model.AdminUserInfo{}
	c.BindJSON(&adminuser)
	totals, currentPage, adminlist, err := ausc.adminuserservice.AddAdminUserInfo(&adminuser)
	return totals, currentPage, adminlist, err
}

//删除管理员信息
func (ausc *AdminUserController) DeleteAdminUserInfoHandler(c *gin.Context) error {
	adminID,_ := strconv.Atoi(c.Query("adminId"))
	err := ausc.adminuserservice.DeleteAdminUserInfo(adminID)
	return err
}