package model

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var UserDB *gorm.DB

func NewDBEngine() {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		"root",
		"",
		"127.0.0.1:3306",
		"gotest",
		"utf8",
		true),
	), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(30)
	if err != nil {
		panic(err)
	}

	UserDB = db
}
