package response_users

import domain_users "github.com/gozzafadillah/27_Docker/praktikum/domain/users"

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
