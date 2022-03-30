package main

import (
	"github.com/gozzafadillah/21_ORM_and_Code_Structure_MVC/praktikum/problem2/config"
	"github.com/gozzafadillah/21_ORM_and_Code_Structure_MVC/praktikum/problem2/routes"
)

func main() {
	config.Init()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
