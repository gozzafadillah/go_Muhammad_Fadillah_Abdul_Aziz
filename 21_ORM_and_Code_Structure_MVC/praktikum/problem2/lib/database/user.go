package database

import (
	"github.com/gozzafadillah/21_ORM_and_Code_Structure_MVC/praktikum/problem2/config"
	"github.com/gozzafadillah/21_ORM_and_Code_Structure_MVC/praktikum/problem2/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}

	return users, nil
}

func GetUser(id int) ([]models.User, error) {
	var user []models.User
	queryData := config.DB.Where("id = ?", id).Find(&user)
	if e := queryData.Error; e != nil {
		return nil, e
	}
	return user, nil
}

func CreateUser(data models.User) (models.User, error) {
	if e := config.DB.Save(&data).Error; e != nil {
		return models.User{}, e
	}
	return data, nil
}

func UpdateUser(id int, data models.User) (models.User, error) {
	queryData := config.DB.Model(&data).Where("id = ?", id).Updates(map[string]interface{}{"id": id, "name": data.Name, "email": data.Email, "password": data.Password})
	if e := queryData.Error; e != nil {
		return models.User{}, e
	}
	return data, nil
}

func DeleteUser(id int) ([]models.User, error) {
	var user []models.User

	queryData := config.DB.Unscoped().Delete(&models.User{}, id)

	if e := queryData.Error; e != nil {
		return nil, e
	}
	return user, nil
}
