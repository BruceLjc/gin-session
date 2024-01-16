package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Nickname       string
	Status         string
	Avatar         string `gorm:"size:1000"`
}

var DB *gorm.DB
var err error

func InitDB() {
	DB, err = gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/singo?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败:" + err.Error())
	}
	fmt.Println("数据库连接成功")
	err = DB.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("映射User表失败:" + err.Error())
	}
	fmt.Println("User表映射成功")
}
