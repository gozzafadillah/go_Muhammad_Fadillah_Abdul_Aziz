package controller

import (
	"net/http"
	"strconv"

	"github.com/gozzafadillah/22_Middleware/praktikum/problem1/lib/database"
	"github.com/gozzafadillah/22_Middleware/praktikum/problem1/models"
	"github.com/labstack/echo/v4"
)

// Controller
func GetBooksController(c echo.Context) error {
	books, e := database.GetBooks()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all book",
		"books":   books,
	})
}

func GetBookController(c echo.Context) error {
	getId, _ := strconv.Atoi(c.Param("id"))
	book, e := database.GetBook(getId)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	if len(book) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Book Not Found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get book by id",
		"book":    book,
	})
}

func CreateBookController(c echo.Context) error {
	temp := models.Book{}
	c.Bind(&temp)
	book, e := database.CreateBook(temp)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success add book",
		"book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	getId, _ := strconv.Atoi(c.Param("id"))

	_, e := database.DeleteBook(getId)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Book",
	})
}

func UpdateBookController(c echo.Context) error {
	// your solution here
	temp := models.Book{}
	c.Bind(&temp)

	getId, _ := strconv.Atoi(c.Param("id"))
	book, e := database.UpdateBook(getId, temp)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Edit Book",
		"book":    book,
	})
}
