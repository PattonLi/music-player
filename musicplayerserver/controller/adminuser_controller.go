package controller

import (
	"fmt"
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"
	//"strconv"
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
func (ausc *AdminUserController) AdminProfileHandler(c *gin.Context) (*model.AdminUserInfo, error) {
	var adminuser *model.AdminUserInfo
	adminuser = nil
	var err error
	fmt.Print(c.Query("adminid"))
	adminname := c.Query("adminname")
	adminuser, err = ausc.adminuserservice.AdminUserInfo(adminname)
	return adminuser, err
}

//获取所有管理员信息
func (ausc *AdminUserController) AllAdminInfoHandler() []gin.H{
	adminlist := ausc.adminuserservice.AllAdminInfo()
	adminusers := make([]gin.H,0)
	for _, adminuser := range adminlist{
		userinfo := gin.H{
			"user_id":    adminuser.ID,
			"username":   adminuser.Adminname,
			"password":   adminuser.Password,
		}
		adminusers = append(adminusers,userinfo)
	}
	return adminusers
}