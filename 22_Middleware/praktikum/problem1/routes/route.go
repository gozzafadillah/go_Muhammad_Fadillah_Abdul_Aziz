package routes

import (
	"github.com/gozzafadillah/22_Middleware/praktikum/problem1/constants"
	"github.com/gozzafadillah/22_Middleware/praktikum/problem1/controller"
	m "github.com/gozzafadillah/22_Middleware/praktikum/problem1/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// Middlewares
	m.LogMiddlewares(e)
	m.TrailingMiddleware(e)

	// create user
	e.POST("/users", controller.CreateUserController)

	// create book
	e.POST("/books", controller.CreateBookController)

	// login
	e.POST("/login", controller.LoginUserController)

	// Basic Auth
	// eAuthBasic := e.Group("/auth")
	// eAuthBasic.Use(middleware.BasicAuth(m.AuthMiddlewareDB))
	// eAuthBasic.GET("/users", controller.GetUsersController)

	// JWT user
	eJWT := e.Group("")
	eJWT.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	eJWT.GET("/users", controller.GetUsersController)
	eJWT.GET("/users/:id", controller.GetUserController)
	eJWT.DELETE("/users/:id", controller.DeleteUserController)
	eJWT.PUT("/users/:id", controller.UpdateUserController)

	// JWT book
	eJWT.GET("/books", controller.GetBooksController)
	eJWT.GET("/books/:id", controller.GetBookController)
	eJWT.DELETE("/books/:id", controller.DeleteBookController)
	eJWT.PUT("/books/:id", controller.UpdateBookController)

	return e
}
