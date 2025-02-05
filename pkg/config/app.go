package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
    d, err := gorm.Open("mysql", "daffa:Axlesharma@12@tcp(192.168.1.100:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }
    db = d
}

func GetDB() *gorm.DB {
	return db
}