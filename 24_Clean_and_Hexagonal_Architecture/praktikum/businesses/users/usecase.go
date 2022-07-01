package business_users

import (
	"errors"

	domain_users "github.com/gozzafadillah/24_Clean_and_Hexagonal_Architecture/praktikum/domain/users"
)

type BusinessUsers struct {
	UserRepo domain_users.Repository
}

func NewUsersBusiness(ur domain_users.Repository) domain_users.Business {
	return BusinessUsers{
		UserRepo: ur,
	}
}

// GetUsers implements domain_users.Business
func (bu BusinessUsers) GetUsers() []domain_users.Users {
	data := bu.UserRepo.GetAllUser()
	return data
}

// Register implements domain_users.Business
func (bu BusinessUsers) Register(domain domain_users.Users) (domain_users.Users, error) {
	data, err := bu.UserRepo.Store(domain)
	if err != nil {
		return domain_users.Users{}, errors.New("create user failed")
	}
	return data, nil
}
