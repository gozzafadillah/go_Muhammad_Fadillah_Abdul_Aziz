package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var (
	users []User
	DB    *gorm.DB
)

// Controller
func GetUsersController(e echo.Context) error {
	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all users",
		"users":   users,
	})
}

func GetUserController(e echo.Context) error {
	getId, _ := strconv.Atoi(e.Param("id"))
	user := User{}

	queryData := DB.Where("id = ?", getId).Find(&user)
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

	if err := DB.Unscoped().Delete(&User{}, ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete User",
	})
}

func UpdateUserController(c echo.Context) error {
	// your solution here
	user := User{}
	c.Bind(&user)

	id, _ := strconv.Atoi(c.Param("id"))

	queryData := DB.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{"id": id, "name": user.Name, "email": user.Email, "password": user.Password})
	if err := queryData.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Edit User",
		"users":   user,
	})
}

func CreateUserController(e echo.Context) error {

	user := User{}
	e.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success add user",
		"user":    user,
	})
}

func init() {
	initDB()
	initMigrate()
}

func initDB() {
	username := "root"
	password := ""
	host := "127.0.0.1"
	port := "3306"
	dbname := "latihan"
	var err error

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&User{})
}

func main() {
	e := echo.New()

	// route
	e.GET("/users", GetUsersController)
	e.GET("/user/:id", GetUserController)
	e.DELETE("/user/:id", DeleteUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/user/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8080"))
}
