package config

import (
	"github.com/gozzafadillah/21_ORM_and_Code_Structure_MVC/praktikum/problem2/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() {
	initDB()
	initMigrate()
}

var DB *gorm.DB

func initDB() {
	username := "root"
	password := ""
	host := "127.0.0.1"
	port := "3306"
	dbname := "latihan"
	var err error

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&model.User{}, &model.Book{})
}
