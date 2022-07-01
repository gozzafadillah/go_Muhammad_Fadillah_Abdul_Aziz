package main

import (
	"github.com/gozzafadillah/24_Clean_and_Hexagonal_Architecture/praktikum/app/config"
	"github.com/gozzafadillah/24_Clean_and_Hexagonal_Architecture/praktikum/app/routes"
	business_users "github.com/gozzafadillah/24_Clean_and_Hexagonal_Architecture/praktikum/businesses/users"
	controller_users "github.com/gozzafadillah/24_Clean_and_Hexagonal_Architecture/praktikum/controllers/users"
	migrate "github.com/gozzafadillah/24_Clean_and_Hexagonal_Architecture/praktikum/migrator"
	mysql_users "github.com/gozzafadillah/24_Clean_and_Hexagonal_Architecture/praktikum/repository/mysql/users"
	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	migrate.AutoMigrate(db)

	e := echo.New()
	userRepo := mysql_users.NewUsersRepo(db)
	userBusiness := business_users.NewUsersBusiness(userRepo)
	userController := controller_users.NewUsersController(userBusiness)

	routeList := routes.RouteList{
		UserBus: userController,
	}

	routeList.NewRouteList(e)
	e.Logger.Fatal(e.Start(":8080"))
}
