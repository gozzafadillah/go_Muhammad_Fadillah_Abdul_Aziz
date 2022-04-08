package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TrailingMiddleware(e *echo.Echo) {
	// Handle ada atau tidak trailing pada endpoint
	e.Pre(middleware.AddTrailingSlash())
	e.Pre(middleware.RemoveTrailingSlash())
}
