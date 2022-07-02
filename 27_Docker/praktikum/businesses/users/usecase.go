package business_users

import (
	"errors"

	middlewares "github.com/gozzafadillah/27_Docker/praktikum/app/middleware"
	domain_users "github.com/gozzafadillah/27_Docker/praktikum/domain/users"
)

type BusinessUsers struct {
	UserRepo domain_users.Repository
	JWTCon   *middlewares.ConfigJwt
}

func NewUsersBusiness(ur domain_users.Repository, jwt *middlewares.ConfigJwt) domain_users.Business {
	return BusinessUsers{
		UserRepo: ur,
		JWTCon:   jwt,
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

// Login implements domain_users.Business
func (bu BusinessUsers) Login(email string, password string) (string, error) {
	data, err := bu.UserRepo.CheckEmailPassword(email, password)
	if err != nil {
		return "", errors.New("email and password miss macth")
	}
	token, err := bu.JWTCon.GenerateToken(data.ID)
	if err != nil {
		return "", errors.New("system cannot generate token")
	}
	return token, nil
}
