package database

import (
	"errors"

	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/config"
	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	queryData := config.DB.Find(&users)
	return users, queryData.Error
}

func GetUser(id int) ([]models.User, error) {
	var user []models.User
	queryData := config.DB.Where("id = ?", id).Find(&user)
	return user, queryData.Error
}

func CreateUser(user models.User) error {
	queryData := config.DB.Save(&user)
	return queryData.Error
}

func UpdateUser(id int, data models.User) (models.User, error) {
	queryData := config.DB.Model(&data).Where("id = ?", id).Updates(map[string]interface{}{"id": id, "name": data.Name, "email": data.Email, "password": data.Password})

	if queryData.RowsAffected == 0 {
		return models.User{}, errors.New("user not found")
	}
	return data, queryData.Error
}

func DeleteUser(id int) error {
	queryData := config.DB.Unscoped().Delete(&models.User{}, id)
	if queryData.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return queryData.Error
}

func LoginUser(data models.User) error {
	queryData := config.DB.Where("email = ? AND password = ?", data.Email, data.Password).First(&data).Error

	return queryData
}
