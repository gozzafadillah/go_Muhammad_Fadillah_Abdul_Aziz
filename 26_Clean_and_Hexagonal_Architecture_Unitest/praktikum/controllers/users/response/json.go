package response_users

import domain_users "github.com/gozzafadillah/26_Clean_and_Hexagonal_Architecture_Unitest/praktikum/domain/users"

type ResponseJSONUsers struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func FromDomainUsers(domain domain_users.Users) ResponseJSONUsers {
	return ResponseJSONUsers{
		Email:    domain.Email,
		Password: domain.Password,
	}
}
