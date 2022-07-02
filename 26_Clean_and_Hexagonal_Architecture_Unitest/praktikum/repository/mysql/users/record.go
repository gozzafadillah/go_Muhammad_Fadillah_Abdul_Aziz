package mysql_users

import (
	domain_users "github.com/gozzafadillah/26_Clean_and_Hexagonal_Architecture_Unitest/praktikum/domain/users"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID       int
	Email    string
	Password string
}

func ToDomainUsers(rec User) domain_users.Users {
	return domain_users.Users{
		ID:       rec.ID,
		Email:    rec.Email,
		Password: rec.Password,
	}
}
