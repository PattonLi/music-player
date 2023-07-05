package dao

import (
	"errors"
	"fmt"
	"music-player/musicplayerserver/model"

	"gorm.io/gorm"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

// 创建表，在第一次运行创建就行了，后面不用管他
func (*UserDao) CreateUserTable() {
	err := DB.AutoMigrate(&model.UserInfo{})
	if err != nil {
		fmt.Println("Create User Table failed: ", err)
	} else {
		fmt.Println("Successfully create User Table.")
	}
}

// 根据用户名或昵称查询用户信息
func (*UserDao) GetUserInfoByName(name string) ([]model.UserInfo, error) {
	var user []model.UserInfo
	err := DB.Find(&user, "username LIKE ? OR nickname LIKE ?", "%"+name+"%", "%"+name+"%").Error
	if errors.Is(err, gorm.ErrRecordNotFound) || DB.RowsAffected == 0 {
		err = errors.New("查找不到用户信息！")
	}
	return user, err
}

// 添加用户
func (*UserDao) AddUser(user *model.UserInfo) (int64, int64, []model.UserInfo,error) {
	newuser := model.UserInfo{}
	var userlist []model.UserInfo
	var totalrecord int64
	var offset int64
	var err error
	var currentPage int64
	DB.Take(&newuser, "phone = ?",user.Phone)
	if(newuser.ID != 0){
		err = errors.New("手机号已注册！")
		return 0, 0, userlist, err
	}
	DB.Create(user)
	DB.Table("user").Count(&totalrecord)
	if(totalrecord % 10 == 0){
		offset = totalrecord - 10
		currentPage = totalrecord / 10
	} else {
		offset = totalrecord - (totalrecord % 10)
		currentPage = totalrecord / 10 + 1
	}
	DB.Offset(int(offset)).Limit(10).Find(&userlist)
	return totalrecord, currentPage, userlist, err
}

// 修改用户信息
func (*UserDao) ModifyUser(user *model.UserInfo) error {
	result := DB.Save(user)
	var err error = nil
	if result.Error != nil {
		err = errors.New("修改失败！")
	}
	return err
}

// 删除用户信息
func (*UserDao) DeleteUser(userID int) error {
	err := DB.Delete(&model.UserInfo{}, userID).Error
	return err
}


// 用户登录验证
func (*UserDao) UserLoginCheck(u *model.UserInfo) (int, error) {
	phone := u.Phone
	password := u.Password
	user := model.UserInfo{}
	err := DB.First(&user, "phone = ? AND password = ?", phone, password).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("用户名不存在或密码错误！")
	}
	return user.ID, err
}

// 手机号验证
func (*UserDao) UserPhoneCheck(u *model.UserInfo) (int, error) {
	phone := u.Phone
	user := model.UserInfo{}
	err := DB.Take(&user, "phone = ?", phone).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查找不到记录！")
	} else {
		err = nil
	}
	return user.ID, err
}

// 返回特定页所有普通用户信息
func (*UserDao) GetAllUserInfo(page int, pagesize int) ([]model.UserInfo, int64) {
	var userlist []model.UserInfo
	var totalrecord int64
	offset := (page-1)*pagesize
	DB.Offset(offset).Limit(pagesize).Find(&userlist).Offset(-1).Limit(-1).Count(&totalrecord)
	return userlist,totalrecord
}

// 根据ID获取单个用户信息
func (*UserDao) GetUserProfile(userID int) (*model.UserInfo, error) {
	user := model.UserInfo{}
	err := DB.First(&user, userID).Error
	return &user, err
}

//更新头像url
func (*UserDao) UpdateUserPicUrl (userID int, url string) error {
	err :=DB.Model(&model.UserInfo{ID: userID}).UpdateColumn("pic_url", url).Error
	if err != nil{
		err = errors.New("更新数据库头像url失败！")
	}
	return err
}
