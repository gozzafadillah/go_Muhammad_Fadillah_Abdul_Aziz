package main

import (
	"github.com/gozzafadillah/22_Middleware/praktikum/problem1/config"
	"github.com/gozzafadillah/22_Middleware/praktikum/problem1/routes"
)

func main() {
	config.Init()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}
