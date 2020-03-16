package models

import (
	"blog/models/user"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {

	var err error
	Db, err = gorm.Open("mysql", "blog:f7fjwF8HpBrbGw2A@(119.23.222.212:3306)/blog?charset=utf8mb4&parseTime=True")

	if nil != err {

		panic(err.Error())
	}
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {

		return "blog_" + defaultTableName
	}
	Db.Set("", "ENGINE=InnoDB").AutoMigrate(
		&user.Members{})

}
