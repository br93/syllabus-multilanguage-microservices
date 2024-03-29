package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

const connection = "root:root4567@tcp(syllabus-mysql:3306)/config_db?charset=utf8&parseTime=True&loc=Local"

func Connect() {
	d, err := gorm.Open(mysql.Open(connection), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
