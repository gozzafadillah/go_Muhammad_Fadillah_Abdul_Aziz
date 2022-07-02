package main

import (
	"github.com/gozzafadillah/27_Docker/praktikum/app/config"
	middlewares "github.com/gozzafadillah/27_Docker/praktikum/app/middleware"
	"github.com/gozzafadillah/27_Docker/praktikum/app/routes"
	business_users "github.com/gozzafadillah/27_Docker/praktikum/businesses/users"
	controller_users "github.com/gozzafadillah/27_Docker/praktikum/controllers/users"
	migrate "github.com/gozzafadillah/27_Docker/praktikum/migrator"
	mysql_users "github.com/gozzafadillah/27_Docker/praktikum/repository/mysql/users"
	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	migrate.AutoMigrate(db)

	configJWT := middlewares.ConfigJwt{
		SecretJWT: "12345",
	}

	e := echo.New()
	userRepo := mysql_users.NewUsersRepo(db)
	userBusiness := business_users.NewUsersBusiness(userRepo, &configJWT)
	userController := controller_users.NewUsersController(userBusiness)

	routeList := routes.RouteList{
		UserBus: userController,
	}

	routeList.NewRouteList(e)
	e.Logger.Fatal(e.Start(":8080"))
}
