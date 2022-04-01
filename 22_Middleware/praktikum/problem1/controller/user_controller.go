package controller

import (
	"net/http"
	"strconv"

	"github.com/gozzafadillah/22_Middleware/praktikum/problem1/config"
	"github.com/gozzafadillah/22_Middleware/praktikum/problem1/lib/database"
	"github.com/gozzafadillah/22_Middleware/praktikum/problem1/middlewares"
	"github.com/gozzafadillah/22_Middleware/praktikum/problem1/models"
	"github.com/labstack/echo/v4"
)

// Controller
func GetUsersController(c echo.Context) error {
	users, _ := database.GetUsers()

	// if e != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	// }

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
	_, e := database.DeleteUser(getId)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete User",
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

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail login",
			"error":   err.Error,
		})
	}

	token, err := middlewares.CreateToken(user.ID, user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail login",
			"error":   err.Error,
		})
	}

	userResponse := models.UserResponse{user.ID, user.Name, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success login",
		"user":    userResponse,
	})
}
