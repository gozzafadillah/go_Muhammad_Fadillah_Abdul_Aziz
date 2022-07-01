package mysql_users

import (
	domain_users "github.com/gozzafadillah/24_Clean_and_Hexagonal_Architecture/praktikum/domain/users"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Email    string
	Password string
}

func ToDomainUsers(rec User) domain_users.Users {
	return domain_users.Users{
		Email:    rec.Email,
		Password: rec.Password,
	}
}
