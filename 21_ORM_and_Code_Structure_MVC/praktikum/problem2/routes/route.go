package routes

import (
	"github.com/gozzafadillah/21_ORM_and_Code_Structure_MVC/praktikum/problem2/controller"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// route
	e.GET("/users", controller.GetUsersController)
	e.GET("/user/:id", controller.GetUserController)
	e.DELETE("/user/:id", controller.DeleteUserController)
	e.POST("/users", controller.CreateUserController)
	e.PUT("/user/:id", controller.UpdateUserController)

	return e
}
