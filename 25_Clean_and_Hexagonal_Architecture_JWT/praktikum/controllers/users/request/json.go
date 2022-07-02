package request_users

import domain_users "github.com/gozzafadillah/25_Clean_and_Hexagonal_Architecture_JWT/praktikum/domain/users"

type RequestJSONUsers struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToDomainUsers(req RequestJSONUsers) domain_users.Users {
	return domain_users.Users{
		Email:    req.Email,
		Password: req.Password,
	}
}
