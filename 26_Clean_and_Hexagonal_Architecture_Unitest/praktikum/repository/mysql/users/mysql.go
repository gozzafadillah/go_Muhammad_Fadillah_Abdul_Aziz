package mysql_users

import (
	domain_users "github.com/gozzafadillah/26_Clean_and_Hexagonal_Architecture_Unitest/praktikum/domain/users"
	"gorm.io/gorm"
)

type UsersRepo struct {
	DB *gorm.DB
}

func NewUsersRepo(db *gorm.DB) domain_users.Repository {
	return UsersRepo{
		DB: db,
	}
}

// GetAllUser implements domain_users.Repository
func (ur UsersRepo) GetAllUser() []domain_users.Users {
	rec := []User{}
	sliceUsers := []domain_users.Users{}
	ur.DB.Find(&rec)
	for _, value := range rec {
		sliceUsers = append(sliceUsers, ToDomainUsers(value))
	}

	return sliceUsers
}

// Store implements domain_users.Repository
func (ur UsersRepo) Store(domain domain_users.Users) (domain_users.Users, error) {
	err := ur.DB.Create(&domain).Error
	return domain, err
}

// CheckEmailPassword implements domain_users.Repository
func (ur UsersRepo) CheckEmailPassword(email string, password string) (domain_users.Users, error) {
	rec := User{}
	err := ur.DB.Where("email = ? AND password = ?", email, password).First(&rec).Error
	return ToDomainUsers(rec), err
}
