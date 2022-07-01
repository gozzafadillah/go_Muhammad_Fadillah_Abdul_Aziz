package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (DB *gorm.DB) {
	var (
		DBNAME = "latihan-clean"
		DBUSER = "root"
		DBPASS = ""
		DBHOST = "127.0.0.1"
		DBPORT = "3306"
	)

	var err error

	dsn := DBUSER + ":" + DBPASS + "@tcp(" + DBHOST + ":" + DBPORT + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return
}
