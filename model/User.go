package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Username string `gorm:"varchar(16);not null"`
	Password string `gorm:"varchar(16);not null"`
	Telphone string `gorm:"varchar(11);not null"`
}
