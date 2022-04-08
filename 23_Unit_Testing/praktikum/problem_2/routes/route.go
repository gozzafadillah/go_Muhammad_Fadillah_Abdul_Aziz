package routes

import (
	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/constants"
	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/controller"
	m "github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// Middlewares
	m.LogMiddlewares(e)
	m.TrailingMiddleware(e)

	// route user
	e.POST("/users", controller.CreateUserController)

	// route book
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)

	// login
	e.POST("/login", controller.LoginUserController)

	// Basic Auth
	// eAuthBasic := e.Group("/auth")
	// eAuthBasic.Use(middleware.BasicAuth(m.AuthMiddlewareDB))
	// eAuthBasic.GET("/users", controller.GetUsersController)

	eJWT := e.Group("")
	eJWT.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// JWT user
	eJWT.GET("/users", controller.GetUsersController)
	eJWT.GET("/users/:id", controller.GetUserController)
	eJWT.DELETE("/users/:id", controller.DeleteUserController)
	eJWT.PUT("/users/:id", controller.UpdateUserController)

	// JWT book
	eJWT.POST("/books", controller.CreateBookController)
	eJWT.DELETE("/books/:id", controller.DeleteBookController)
	eJWT.PUT("/books/:id", controller.UpdateBookController)

	return e
}
