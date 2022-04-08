package controller

import (
	"net/http"
	"strconv"

	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/lib/database"
	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/models"
	"github.com/labstack/echo/v4"
)

// Controller
func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Database Broke",
		})
	} else if len(books) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Database Empty",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all book",
		"books":   books,
	})
}

func GetBookController(c echo.Context) error {
	getId, _ := strconv.Atoi(c.Param("id"))
	book, err := database.GetBook(getId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Database Broke",
		})
	} else if len(book) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Book not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get book",
		"book":    book,
	})
}

var SaveBook = func(book models.Book) error {
	return database.CreateBook(book)
}

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	err := SaveBook(book)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed add book",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success add book",
		"book":    book,
	})
}

var DeleteBook = func(id int) error {
	return database.DeleteBook(id)
}

func DeleteBookController(c echo.Context) error {
	getId, _ := strconv.Atoi(c.Param("id"))
	err := DeleteBook(getId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Book",
	})
}

var UpdateBook = func(id int, data models.Book) error {
	return database.UpdateBook(id, data)
}

func UpdateBookController(c echo.Context) error {
	// your solution here
	book := models.Book{}
	c.Bind(&book)

	getId, _ := strconv.Atoi(c.Param("id"))
	err := UpdateBook(getId, book)

	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Failed update data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Edit Book",
		"book":    book,
	})
}
