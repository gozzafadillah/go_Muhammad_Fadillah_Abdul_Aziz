package controller

import (
	"net/http"
	"strconv"

	"github.com/gozzafadillah/21_ORM_and_Code_Structure_MVC/praktikum/problem2/config"
	"github.com/gozzafadillah/21_ORM_and_Code_Structure_MVC/praktikum/problem2/model"
	"github.com/labstack/echo/v4"
)

var (
	books []model.Book
)

// Controller
func GetBooksController(e echo.Context) error {
	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all book",
		"books":   books,
	})
}

func GetBookController(e echo.Context) error {
	getId, _ := strconv.Atoi(e.Param("id"))
	book := model.Book{}

	queryData := config.DB.Where("id = ?", getId).Find(&book)
	if err := queryData.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get book by id",
		"book":    book,
	})
}

func DeleteBookController(e echo.Context) error {
	getId, _ := strconv.Atoi(e.Param("id"))

	if err := config.DB.Unscoped().Delete(&model.Book{}, getId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Book",
	})
}

func UpdateBookController(c echo.Context) error {
	// your solution here
	book := model.Book{}
	c.Bind(&book)

	id, _ := strconv.Atoi(c.Param("id"))

	queryData := config.DB.Model(&book).Where("id = ?", id).Updates(map[string]interface{}{"id": id, "title": book.Title, "author": book.Author, "publisher": book.Publisher})
	if err := queryData.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Edit Book",
		"book":    book,
	})
}

func CreateBookController(e echo.Context) error {

	book := model.Book{}
	e.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success add user",
		"book":    book,
	})
}
