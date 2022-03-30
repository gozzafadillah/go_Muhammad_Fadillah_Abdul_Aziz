package controller

import (
	"net/http"
	"strconv"

	"github.com/gozzafadillah/21_ORM_and_Code_Structure_MVC/praktikum/problem2/config"
	"github.com/gozzafadillah/21_ORM_and_Code_Structure_MVC/praktikum/problem2/model"
	"github.com/labstack/echo/v4"
)

var (
	users []model.User
)

// Controller
func GetUsersController(e echo.Context) error {
	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all users",
		"users":   users,
	})
}

func GetUserController(e echo.Context) error {
	getId, _ := strconv.Atoi(e.Param("id"))
	user := model.User{}

	queryData := config.DB.Where("id = ?", getId).Find(&user)
	if err := queryData.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get user",
		"user":    user,
	})
}

func DeleteUserController(e echo.Context) error {
	ID, _ := strconv.Atoi(e.Param("id"))

	if err := config.DB.Unscoped().Delete(&model.User{}, ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete User",
	})
}

func UpdateUserController(c echo.Context) error {
	// your solution here
	user := model.User{}
	c.Bind(&user)

	id, _ := strconv.Atoi(c.Param("id"))

	queryData := config.DB.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{"id": id, "name": user.Name, "email": user.Email, "password": user.Password})
	if err := queryData.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Edit User",
		"users":   user,
	})
}

func CreateUserController(e echo.Context) error {

	user := model.User{}
	e.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success add user",
		"user":    user,
	})
}
