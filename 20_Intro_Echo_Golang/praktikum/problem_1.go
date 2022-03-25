package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// Controller
func GetUsersController(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all users",
		"users":   users,
	})
}

func GetUserController(e echo.Context) error {
	getId, _ := strconv.Atoi(e.Param("id"))
	user := []interface{}{}

	for i := 0; i <= len(users)-1; i++ {
		if users[i].Id == getId {
			user = append(user, users[i])
		}
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get user",
		"user":    user,
	})
}

func DeleteUserController(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	for i := 0; i <= len(users)-1; i++ {
		if id == users[i].Id {
			users = append(users[:i], users[i+1:]...)
		}
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete User",
		"users":   users,
	})

}

func UpdateUserController(c echo.Context) error {
	// your solution here
	user := User{}
	c.Bind(&user)

	id, _ := strconv.Atoi(c.Param("id"))

	for i := 0; i <= len(users)-1; i++ {
		if id == users[i].Id {
			users[i].Name = user.Name
			users[i].Password = user.Password
			users[i].Id = user.Id
			users[i].Email = user.Email
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Edit User",
		"users":   users[id],
	})
}

func CreateUserController(e echo.Context) error {

	user := User{}
	e.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success add user",
		"user":    user,
	})
}

func main() {
	e := echo.New()

	// route
	e.GET("/users", GetUsersController)
	e.GET("/user/:id", GetUserController)
	e.DELETE("/user/:id", DeleteUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/user/:id", UpdateUserController)

	e.Start(":8080")
}
