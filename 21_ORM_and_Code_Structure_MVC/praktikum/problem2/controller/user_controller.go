package controller

import (
	"net/http"
	"strconv"

	"github.com/gozzafadillah/21_ORM_and_Code_Structure_MVC/praktikum/problem2/lib/database"
	"github.com/gozzafadillah/21_ORM_and_Code_Structure_MVC/praktikum/problem2/models"
	"github.com/labstack/echo/v4"
)

// Controller
func GetUsersController(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all users",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	getId, _ := strconv.Atoi(c.Param("id"))
	user, e := database.GetUser(getId)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	if len(user) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "User not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get user",
		"user":    user,
	})

}

func CreateUserController(c echo.Context) error {
	temp := models.User{}

	c.Bind(&temp)

	user, e := database.CreateUser(temp)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success add user",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	getId, _ := strconv.Atoi(c.Param("id"))
	user, e := database.DeleteUser(getId)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete User",
		"user":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	temp := models.User{}
	getId, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&temp)
	user, e := database.UpdateUser(getId, temp)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Edit User",
		"users":   user,
	})
}
