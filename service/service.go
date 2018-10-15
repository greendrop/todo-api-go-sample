package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func InitDB() {
	user := "root"
	password := "mysql"
	dbname := "todo"

	db, err = gorm.Open("mysql", user+":"+password+"@tcp(mysql:3306)/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
}

func GetDBConnection() *gorm.DB {
	return db
}
