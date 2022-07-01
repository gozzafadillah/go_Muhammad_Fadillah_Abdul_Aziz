package main

import (
	"github.com/gozzafadillah/24_Clean_and_Hexagonal_Architecture/praktikum/belajar-go-echo/config"
	"github.com/gozzafadillah/24_Clean_and_Hexagonal_Architecture/praktikum/belajar-go-echo/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	app := echo.New()
	app.GET("/users", controller.GetAllUsers(db))
	app.POST("/users", controller.CreateUser(db))
	app.Start(":8080")
}
