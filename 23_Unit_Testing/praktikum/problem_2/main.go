package main

import (
	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/config"
	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/routes"
)

func main() {
	config.Init()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}
