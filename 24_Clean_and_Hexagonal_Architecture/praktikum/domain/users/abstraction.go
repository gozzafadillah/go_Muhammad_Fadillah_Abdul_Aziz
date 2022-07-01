package domain_users

type Business interface {
	GetUsers() []Users
	Register(domain Users) (Users, error)
}

type Repository interface {
	GetAllUser() []Users
	Store(domain Users) (Users, error)
}
