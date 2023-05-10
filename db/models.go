package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB  *gorm.DB
	ERR error
)

func init() {
	DB, ERR = gorm.Open("sqlite3", "test.db")
	if ERR != nil {
		fmt.Println(ERR)
	}
	DB.LogMode(true)
	DB.SingularTable(true)
	DB.Table("user").CreateTable(&User{})       // 创建用户表
	DB.Table("article").CreateTable(&Article{}) // 创建文章表
	DB.Table("role").CreateTable(&Role{})
	DB.Table("role_with_user").CreateTable(&RoleWithUser{})
	DB.Table("article_permission").CreateTable(&ArticlePermission{})
}
