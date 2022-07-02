package domain_users

type Business interface {
	GetUsers() []Users
	Register(domain Users) (Users, error)
	Login(email, password string) (string, error)
}

type Repository interface {
	GetAllUser() []Users
	Store(domain Users) (Users, error)
	CheckEmailPassword(email, password string) (Users, error)
}
