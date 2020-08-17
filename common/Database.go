package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-demo/model"
)

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "root"
	username := "root"
	password := "root"
	charset := "utf8"
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, url)
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&model.User{})
	return db
}
