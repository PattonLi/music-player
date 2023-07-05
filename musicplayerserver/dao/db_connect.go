package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库操作对象DB，导出的首字母要大写
var DB *gorm.DB

func Init() {
	// DSN stands for "Data Source Name". It is a string that contains the necessary information to establish a database connection
	dsn := "lpt:Lpt123456.@tcp(47.120.38.64:3306)/music?charset=utf8mb4&parseTime=True&loc=Local"
	// 打开连接数据库，默认配置
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Open Database failed: ", err)
	} else {
		fmt.Println("Successfully open Database.")
	}
	DB = db
}
