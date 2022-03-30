package database

import (
	"errors"

	"github.com/gozzafadillah/21_ORM_and_Code_Structure_MVC/praktikum/problem2/config"
	"github.com/gozzafadillah/21_ORM_and_Code_Structure_MVC/praktikum/problem2/models"
)

func GetBooks() ([]models.Book, error) {
	var books []models.Book
	queryData := config.DB.Find(&books)
	if e := queryData.Error; e != nil {
		return nil, e
	}
	return books, nil
}

func GetBook(id int) ([]models.Book, error) {
	var book []models.Book
	queryData := config.DB.Where("id = ?", id).Find(&book)
	if e := queryData.Error; e != nil {
		return nil, e
	}
	return book, nil
}

func CreateBook(data models.Book) (models.Book, error) {
	queryData := config.DB.Save(&data)
	if e := queryData.Error; e != nil {
		return models.Book{}, e
	}
	return data, nil
}

func UpdateBook(id int, data models.Book) (models.Book, error) {
	queryData := config.DB.Model(&data).Where("id = ?", id).Updates(map[string]interface{}{"id": id, "title": data.Title, "author": data.Author, "publisher": data.Publisher})
	if queryData.RowsAffected == 0 {
		return models.Book{}, errors.New("no data updated")
	}
	if e := queryData.Error; e != nil {
		return models.Book{}, e
	}
	return data, nil
}

func DeleteBook(id int) ([]models.Book, error) {
	var book []models.Book
	queryData := config.DB.Unscoped().Where("id = ?", id).Delete(&book)
	if e := queryData.Error; e != nil {
		return nil, e
	}
	return book, nil
}
