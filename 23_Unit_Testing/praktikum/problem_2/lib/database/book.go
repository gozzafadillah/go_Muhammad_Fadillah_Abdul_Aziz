package database

import (
	"errors"

	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/config"
	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/models"
)

func GetBooks() ([]models.Book, error) {
	var books []models.Book
	queryData := config.DB.Find(&books)
	return books, queryData.Error
}

func GetBook(id int) ([]models.Book, error) {
	var book []models.Book
	queryData := config.DB.Where("id = ?", id).Find(&book)
	return book, queryData.Error
}

func CreateBook(data models.Book) error {
	queryData := config.DB.Save(&data)
	return queryData.Error
}

func UpdateBook(id int, data models.Book) error {
	queryData := config.DB.Model(&data).Where("id = ?", id).Updates(map[string]interface{}{"id": id, "title": data.Title, "author": data.Author, "publisher": data.Publisher})

	return queryData.Error
}

func DeleteBook(id int) error {
	var book []models.Book
	queryData := config.DB.Unscoped().Where("id = ?", id).Delete(&book)
	if queryData.RowsAffected == 0 {
		return errors.New("book not found")
	}

	return queryData.Error
}
