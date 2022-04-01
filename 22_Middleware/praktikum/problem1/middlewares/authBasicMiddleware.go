package middlewares

import (
	"github.com/gozzafadillah/22_Middleware/praktikum/problem1/config"
	"github.com/gozzafadillah/22_Middleware/praktikum/problem1/models"
	"github.com/labstack/echo/v4"
)

func AuthMiddlewareDB(email, password string, c echo.Context) (bool, error) {
	var user models.User
	err := config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
