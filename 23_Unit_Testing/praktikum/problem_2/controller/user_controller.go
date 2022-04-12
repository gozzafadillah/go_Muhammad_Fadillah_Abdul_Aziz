package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/lib/database"
	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/middlewares"
	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/models"
	"github.com/labstack/echo/v4"
)

// Controller
func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Database Broke",
		})
	} else if len(users) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Database empty",
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success get all users",
			"users":   users,
		})
	}

}

func GetUserController(c echo.Context) error {
	getId, _ := strconv.Atoi(c.Param("id"))
	user, err := database.GetUser(getId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Database Broke",
		})
	} else if len(user) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "User not found",
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success get user",
			"user":    user,
		})
	}

}

var UserSave = func(user models.User) error {
	return database.CreateUser(user)
}

func CreateUserController(c echo.Context) error {
	user := models.User{}

	c.Bind(&user)

	if err := UserSave(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed save data !",
			"user":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success add user",
		"user":    user,
	})
}

var DeleteUser = func(id int) error {
	return database.DeleteUser(id)
}

func DeleteUserController(c echo.Context) error {
	getId, _ := strconv.Atoi(c.Param("id"))
	err := DeleteUser(getId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Delete",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete User",
	})
}

func UpdateUserController(c echo.Context) error {
	temp := models.User{}
	getId, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&temp)
	user, err := database.UpdateUser(getId, temp)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Edit User",
		"users":   user,
	})
}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	err := database.LoginUser(user)
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Fail login",
			"error":   err.Error(),
		})
	}

	token, err := middlewares.CreateToken(user.ID, user.Name)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Fail login",
			"error":   err.Error,
		})
	}

	userResponse := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success login",
		"user":    userResponse,
	})
}
